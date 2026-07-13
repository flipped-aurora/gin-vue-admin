package example

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func TestScanUploadTokenIsBoundAndOneTime(t *testing.T) {
	originalUseRedis := global.GVA_CONFIG.System.UseRedis
	defer func() { global.GVA_CONFIG.System.UseRedis = originalUseRedis }()
	global.GVA_CONFIG.System.UseRedis = false
	store := newScanUploadTokenStore()
	now := time.Date(2026, time.July, 13, 0, 0, 0, 0, time.UTC)
	store.now = func() time.Time { return now }

	token, expiresAt, err := store.issue(42)
	if err != nil {
		t.Fatalf("issue token: %v", err)
	}
	if token == "" {
		t.Fatal("expected a non-empty token")
	}
	if expiresAt.Sub(now) != scanUploadTokenTTL {
		t.Fatalf("unexpected expiry: %v", expiresAt)
	}

	classID, err := store.consume(token)
	if err != nil {
		t.Fatalf("consume token: %v", err)
	}
	if classID != 42 {
		t.Fatalf("expected class ID 42, got %d", classID)
	}
	if _, err = store.consume(token); !errors.Is(err, ErrInvalidScanUploadToken) {
		t.Fatalf("expected token reuse to fail, got %v", err)
	}
}

func TestScanUploadTokenExpires(t *testing.T) {
	originalUseRedis := global.GVA_CONFIG.System.UseRedis
	defer func() { global.GVA_CONFIG.System.UseRedis = originalUseRedis }()
	global.GVA_CONFIG.System.UseRedis = false
	store := newScanUploadTokenStore()
	now := time.Date(2026, time.July, 13, 0, 0, 0, 0, time.UTC)
	store.now = func() time.Time { return now }

	token, _, err := store.issue(7)
	if err != nil {
		t.Fatalf("issue token: %v", err)
	}
	now = now.Add(scanUploadTokenTTL)
	if _, err = store.consume(token); !errors.Is(err, ErrInvalidScanUploadToken) {
		t.Fatalf("expected expired token to fail, got %v", err)
	}
}

func TestScanUploadTokenConcurrentConsumeSucceedsOnce(t *testing.T) {
	originalUseRedis := global.GVA_CONFIG.System.UseRedis
	defer func() { global.GVA_CONFIG.System.UseRedis = originalUseRedis }()
	global.GVA_CONFIG.System.UseRedis = false
	store := newScanUploadTokenStore()
	token, _, err := store.issue(9)
	if err != nil {
		t.Fatalf("issue token: %v", err)
	}

	var successes atomic.Int32
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if classID, consumeErr := store.consume(token); consumeErr == nil && classID == 9 {
				successes.Add(1)
			}
		}()
	}
	wg.Wait()
	if successes.Load() != 1 {
		t.Fatalf("expected exactly one successful consume, got %d", successes.Load())
	}
}
