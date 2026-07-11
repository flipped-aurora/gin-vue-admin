package sse

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Event 一条 SSE 推送事件。Name 对应前端 EventSource 的事件名
// (空或 "message" 触发 onmessage,自定义名走 addEventListener),
// Data 建议为 JSON 字符串;多行内容会按 SSE 规范拆成多个 data: 行。
type Event struct {
	ID   string
	Name string
	Data string
}

// Subscriber 一个活跃的下游连接。同一用户可同时持有多个(多标签页/多设备)。
type Subscriber struct {
	userID    uint
	ch        chan Event
	done      chan struct{} // 关闭即通知 Stream 退出(逐出/停机)
	closeOnce sync.Once
}

func (s *Subscriber) close() {
	s.closeOnce.Do(func() { close(s.done) })
}

// subscriberBuf 单连接事件缓冲:满则丢弃本次事件(慢消费者不阻塞发布方),
// 客户端靠下一次事件或重连后的全量拉取补齐状态。
const subscriberBuf = 16

// maxSubsPerUser 单用户最大并发连接数:超出时逐出最早注册的连接,
// 防止单账号低成本囤积长连接耗尽服务端资源。
const maxSubsPerUser = 10

// Hub 按用户维度管理 SSE 长连接的进程内推送中枢。
// 这是 GVA 本体的通用地基:任何需要实时推送的功能(消息中心/告警/在线状态)
// 都应复用它,而不是各自管理连接。
// 注意:单实例进程内实现;多实例部署时需在上层用共享队列(Redis 等)把事件
// 扇出到每个实例的 Hub,本包刻意不引入外部依赖。
type Hub struct {
	mu   sync.RWMutex
	subs map[uint]map[*Subscriber]struct{}
}

func NewHub() *Hub {
	return &Hub{subs: make(map[uint]map[*Subscriber]struct{})}
}

var defaultHub = NewHub()

// Default 全局默认 Hub,业务方直接使用它即可互通(同一进程同一连接池)
func Default() *Hub { return defaultHub }

// Subscribe 为用户注册一个连接,返回订阅句柄;用完必须 Unsubscribe。
// 用户连接数达到上限时逐出一个既有连接(其 Stream 循环随即退出)。
func (h *Hub) Subscribe(userID uint) *Subscriber {
	s := &Subscriber{userID: userID, ch: make(chan Event, subscriberBuf), done: make(chan struct{})}
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.subs[userID] == nil {
		h.subs[userID] = make(map[*Subscriber]struct{})
	}
	if len(h.subs[userID]) >= maxSubsPerUser {
		for old := range h.subs[userID] {
			delete(h.subs[userID], old)
			old.close()
			break
		}
	}
	h.subs[userID][s] = struct{}{}
	return s
}

// Unsubscribe 注销连接并清理空桶。幂等。
func (h *Hub) Unsubscribe(s *Subscriber) {
	if s == nil {
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	if set, ok := h.subs[s.userID]; ok {
		delete(set, s)
		if len(set) == 0 {
			delete(h.subs, s.userID)
		}
	}
}

// Publish 向指定用户的所有活跃连接投递事件。
// 用户不在线时静默丢弃——离线消息由业务落库、上线后拉取,Hub 只负责"在线加速"。
// 投递非阻塞:单连接缓冲满时丢弃该连接的本次事件。
func (h *Hub) Publish(userID uint, e Event) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for s := range h.subs[userID] {
		select {
		case s.ch <- e:
		default: // 慢消费者:丢弃,不阻塞发布方
		}
	}
}

// PublishToUsers 批量投递(只对在线用户生效)
func (h *Hub) PublishToUsers(userIDs []uint, e Event) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for _, uid := range userIDs {
		for s := range h.subs[uid] {
			select {
			case s.ch <- e:
			default:
			}
		}
	}
}

// Broadcast 向所有在线连接投递事件
func (h *Hub) Broadcast(e Event) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for _, set := range h.subs {
		for s := range set {
			select {
			case s.ch <- e:
			default:
			}
		}
	}
}

// Shutdown 关闭全部活跃连接(通知各 Stream 循环退出)并清空注册表。
// 必须在 http.Server 优雅停机(srv.Shutdown)之前调用:SSE 常驻连接
// 永不空闲,不先关闭它们,Shutdown 会等到超时。
func (h *Hub) Shutdown() {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, set := range h.subs {
		for s := range set {
			s.close()
		}
	}
	h.subs = make(map[uint]map[*Subscriber]struct{})
}

// Online 用户是否有活跃连接
func (h *Hub) Online(userID uint) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.subs[userID]) > 0
}

// OnlineCount 在线用户数(按用户去重,非连接数)
func (h *Hub) OnlineCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.subs)
}

// Stream 以 SSE 协议服务一个 gin 请求,阻塞直到客户端断开:
// 设置流式响应头、解除本连接的服务器写超时、注册订阅、周期心跳保活。
// heartbeat<=0 时默认 30s。调用方负责鉴权并传入可信 userID。
// 注意:挂载该 handler 的路由绝不能套 TimeoutMiddleware(其 goroutine+
// AbortWithStatusJSON 模型与流式响应冲突)。
func (h *Hub) Stream(c *gin.Context, userID uint, heartbeat time.Duration) {
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.String(http.StatusInternalServerError, "streaming unsupported")
		return
	}

	header := c.Writer.Header()
	header.Set("Content-Type", "text/event-stream; charset=utf-8")
	header.Set("Cache-Control", "no-cache, no-transform")
	header.Set("Connection", "keep-alive")
	header.Set("X-Accel-Buffering", "no") // 关闭 nginx 缓冲

	// 解除本连接的读/写超时:core 的 http.Server Read/WriteTimeout(当前 10min)
	// 都会掐断常驻连接——写超时直接断写,读超时到期会让服务器的后台读失败进而
	// 取消请求 ctx。gin 的 responseWriter 实现了 Unwrap,ResponseController 可
	// 穿透到原始 writer;失败则容忍,由 EventSource 自动重连兜底。
	rc := http.NewResponseController(c.Writer)
	_ = rc.SetWriteDeadline(time.Time{})
	_ = rc.SetReadDeadline(time.Time{})

	sub := h.Subscribe(userID)
	defer h.Unsubscribe(sub)

	c.Status(http.StatusOK)
	// 首个注释帧:确认连接建立并穿透中间层缓冲
	_, _ = c.Writer.WriteString(": connected\n\n")
	flusher.Flush()

	if heartbeat <= 0 {
		heartbeat = 30 * time.Second
	}
	tick := time.NewTicker(heartbeat)
	defer tick.Stop()

	ctx := c.Request.Context()
	for {
		select {
		case <-ctx.Done():
			return
		case <-sub.done: // 被逐出或 Hub 停机
			return
		case e := <-sub.ch:
			writeEvent(c.Writer, e)
			flusher.Flush()
		case <-tick.C:
			// SSE 注释行心跳:保活代理/防 idle 断连,客户端不可见
			_, _ = c.Writer.WriteString(": ping\n\n")
			flusher.Flush()
		}
	}
}

// sanitizeField 清洗 id/event 字段:换行/回车会破坏 SSE 帧结构(帧注入),
// 作为通用地基对不可信输入做防御性剥离
func sanitizeField(s string) string {
	if !strings.ContainsAny(s, "\r\n") {
		return s
	}
	return strings.NewReplacer("\r", "", "\n", "").Replace(s)
}

// writeEvent 按 SSE 规范序列化事件(id/event 行清洗换行,多行 data 拆分)
func writeEvent(w gin.ResponseWriter, e Event) {
	if e.ID != "" {
		_, _ = w.WriteString("id: " + sanitizeField(e.ID) + "\n")
	}
	if e.Name != "" && e.Name != "message" {
		_, _ = w.WriteString("event: " + sanitizeField(e.Name) + "\n")
	}
	for line := range strings.SplitSeq(e.Data, "\n") {
		_, _ = w.WriteString("data: " + strings.TrimSuffix(line, "\r") + "\n")
	}
	_, _ = w.WriteString("\n")
}
