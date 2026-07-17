package sse

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestHubSubscribePublishUnsubscribe(t *testing.T) {
	h := NewHub()
	s1 := h.Subscribe(7)
	s2 := h.Subscribe(7) // 同一用户第二个连接(多标签页)
	other := h.Subscribe(8)

	if !h.Online(7) || h.OnlineCount() != 2 {
		t.Fatalf("online state wrong: online(7)=%v count=%d", h.Online(7), h.OnlineCount())
	}

	h.Publish(7, Event{Name: "message", Data: "hello"})
	for i, s := range []*Subscriber{s1, s2} {
		select {
		case e := <-s.ch:
			if e.Data != "hello" {
				t.Fatalf("sub%d got wrong event: %+v", i+1, e)
			}
		default:
			t.Fatalf("sub%d should receive event", i+1)
		}
	}
	select {
	case <-other.ch:
		t.Fatal("other user must not receive")
	default:
	}

	h.Unsubscribe(s1)
	h.Unsubscribe(s2)
	if h.Online(7) {
		t.Fatal("user 7 should be offline after unsubscribe")
	}
	h.Unsubscribe(s1) // 幂等
}

func TestHubSlowConsumerDoesNotBlock(t *testing.T) {
	h := NewHub()
	s := h.Subscribe(1)
	done := make(chan struct{})
	go func() {
		// 缓冲 16,发 100 条:满后丢弃,不得阻塞
		for i := 0; i < 100; i++ {
			h.Publish(1, Event{Data: "x"})
		}
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("publish blocked by slow consumer")
	}
	if got := len(s.ch); got != subscriberBuf {
		t.Fatalf("buffer should hold %d events, got %d", subscriberBuf, got)
	}
}

func TestHubPerUserCapEvictsOldest(t *testing.T) {
	h := NewHub()
	subs := make([]*Subscriber, 0, maxSubsPerUser+1)
	for i := 0; i < maxSubsPerUser+1; i++ {
		subs = append(subs, h.Subscribe(5))
	}
	h.mu.RLock()
	got := len(h.subs[5])
	h.mu.RUnlock()
	if got != maxSubsPerUser {
		t.Fatalf("per-user connections should cap at %d, got %d", maxSubsPerUser, got)
	}
	// 恰有一个既有连接被逐出(done 已关闭)
	evicted := 0
	for _, s := range subs {
		select {
		case <-s.done:
			evicted++
		default:
		}
	}
	if evicted != 1 {
		t.Fatalf("exactly one subscriber should be evicted, got %d", evicted)
	}
}

func TestHubShutdownClosesAll(t *testing.T) {
	h := NewHub()
	a, b := h.Subscribe(1), h.Subscribe(2)
	h.Shutdown()
	for i, s := range []*Subscriber{a, b} {
		select {
		case <-s.done:
		default:
			t.Fatalf("subscriber %d should be closed after Shutdown", i)
		}
	}
	if h.OnlineCount() != 0 {
		t.Fatal("registry should be empty after Shutdown")
	}
	// 停机后 Publish 不应 panic
	h.Publish(1, Event{Data: "x"})
}

func TestWriteEventSanitizesFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	writeEvent(c.Writer, Event{ID: "1\nid: 2", Name: "evil\nevent: fake", Data: "a\r\nb"})
	out := rec.Body.String()
	// 帧注入判定按"行":换行剥离后注入内容只能拼在原行里,不得成为独立的帧字段行
	for line := range strings.SplitSeq(out, "\n") {
		if line == "id: 2" || line == "event: fake" {
			t.Fatalf("frame injection not sanitized: %q", out)
		}
	}
	if !strings.Contains(out, "data: a\ndata: b\n") {
		t.Fatalf("CRLF data should split cleanly: %q", out)
	}
}

func TestHubBroadcastAndPublishToUsers(t *testing.T) {
	h := NewHub()
	a, b := h.Subscribe(1), h.Subscribe(2)
	h.Broadcast(Event{Data: "all"})
	if len(a.ch) != 1 || len(b.ch) != 1 {
		t.Fatal("broadcast should reach every subscriber")
	}
	h.PublishToUsers([]uint{2}, Event{Data: "only2"})
	if len(a.ch) != 1 || len(b.ch) != 2 {
		t.Fatal("publishToUsers should only reach listed users")
	}
}

func TestStreamWritesEventsAndCleansUp(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewHub()

	r := gin.New()
	r.GET("/stream", func(c *gin.Context) { h.Stream(c, 9, time.Hour) })

	srv := httptest.NewServer(r)
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/stream")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if ct := resp.Header.Get("Content-Type"); !strings.Contains(ct, "text/event-stream") {
		t.Fatalf("wrong content type: %q", ct)
	}

	reader := bufio.NewReader(resp.Body)
	// 首帧 ": connected"
	line, _ := reader.ReadString('\n')
	if !strings.HasPrefix(line, ": connected") {
		t.Fatalf("expected connected comment, got %q", line)
	}

	// 等订阅注册完成后投递
	deadline := time.Now().Add(2 * time.Second)
	for !h.Online(9) {
		if time.Now().After(deadline) {
			t.Fatal("subscriber not registered")
		}
		time.Sleep(10 * time.Millisecond)
	}
	h.Publish(9, Event{ID: "1", Name: "notify", Data: "line1\nline2"})

	var got []string
	for i := 0; i < 6; i++ {
		l, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		got = append(got, strings.TrimRight(l, "\n"))
		if l == "\n" && len(got) > 1 {
			break
		}
	}
	joined := strings.Join(got, "|")
	for _, want := range []string{"id: 1", "event: notify", "data: line1", "data: line2"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("stream output missing %q, got %q", want, joined)
		}
	}

	// 断开后订阅应被清理
	resp.Body.Close()
	deadline = time.Now().Add(2 * time.Second)
	for h.Online(9) {
		if time.Now().After(deadline) {
			t.Fatal("subscriber not cleaned up after disconnect")
		}
		time.Sleep(10 * time.Millisecond)
	}
}
