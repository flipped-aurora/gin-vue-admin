// server/task/registry_test.go
package task

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"testing"
)

func TestRegistryRegisterGetList(t *testing.T) {
	Register("t_b", "任务B", func(ctx context.Context, _ json.RawMessage) error { return nil })
	Register("t_a", "任务A", func(ctx context.Context, _ json.RawMessage) error { return errors.New("boom") })

	fn, ok := Get("t_a")
	if !ok {
		t.Fatal("t_a 应已注册")
	}
	if err := fn(context.Background(), nil); err == nil || err.Error() != "boom" {
		t.Fatalf("期望 boom, got %v", err)
	}
	if _, ok := Get("not_exist"); ok {
		t.Fatal("不存在的任务不应命中")
	}

	// List 稳定排序且包含元信息
	metas := List()
	idxA, idxB := -1, -1
	for i, m := range metas {
		if m.Name == "t_a" {
			idxA = i
			if m.Description != "任务A" {
				t.Fatalf("描述不符: %s", m.Description)
			}
		}
		if m.Name == "t_b" {
			idxB = i
		}
	}
	if idxA == -1 || idxB == -1 || idxA > idxB {
		t.Fatalf("List 应含 t_a/t_b 且按名排序, idxA=%d idxB=%d", idxA, idxB)
	}
}

func TestRegistryOverwrite(t *testing.T) {
	Register("t_ow", "v1", func(ctx context.Context, _ json.RawMessage) error { return errors.New("v1") })
	Register("t_ow", "v2", func(ctx context.Context, _ json.RawMessage) error { return errors.New("v2") })
	fn, _ := Get("t_ow")
	if err := fn(context.Background(), nil); err.Error() != "v2" {
		t.Fatal("同名重复注册应以最后一次为准(reload 幂等)")
	}
}

func TestRegistryConcurrent(t *testing.T) { // 配合 -race
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func(n int) {
			defer wg.Done()
			Register(fmt.Sprintf("c_%d", n%10), "并发", func(ctx context.Context, _ json.RawMessage) error { return nil })
		}(i)
		go func() {
			defer wg.Done()
			_ = List()
			_, _ = Get("c_1")
		}()
	}
	wg.Wait()
}
