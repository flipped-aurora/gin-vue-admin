package example

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/redis/go-redis/v9"
)

const (
	scanUploadTokenTTL    = 5 * time.Minute
	scanUploadTokenPrefix = "scan-upload-token:"
)

var (
	ErrInvalidScanUploadToken = errors.New("扫码上传凭证无效或已过期")
	scanUploadTokens          = newScanUploadTokenStore()
	consumeScanUploadToken    = redis.NewScript(`
local value = redis.call("GET", KEYS[1])
if value then
  redis.call("DEL", KEYS[1])
end
return value
`)
)

type scanUploadTokenEntry struct {
	classID   int
	expiresAt time.Time
}

type scanUploadTokenStore struct {
	mu      sync.Mutex
	entries map[string]scanUploadTokenEntry
	now     func() time.Time
}

func newScanUploadTokenStore() *scanUploadTokenStore {
	return &scanUploadTokenStore{
		entries: make(map[string]scanUploadTokenEntry),
		now:     time.Now,
	}
}

func newScanUploadToken() (string, error) {
	randomBytes := make([]byte, 32)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("生成扫码上传凭证失败: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(randomBytes), nil
}

func (s *scanUploadTokenStore) issue(classID int) (string, time.Time, error) {
	for {
		token, err := newScanUploadToken()
		if err != nil {
			return "", time.Time{}, err
		}
		expiresAt := s.now().Add(scanUploadTokenTTL)

		if global.GVA_CONFIG.System.UseRedis && global.GVA_REDIS != nil {
			created, err := global.GVA_REDIS.SetNX(
				context.Background(),
				scanUploadTokenPrefix+token,
				strconv.Itoa(classID),
				scanUploadTokenTTL,
			).Result()
			if err != nil {
				return "", time.Time{}, fmt.Errorf("保存扫码上传凭证失败: %w", err)
			}
			if created {
				return token, expiresAt, nil
			}
			continue
		}

		s.mu.Lock()
		s.removeExpiredLocked()
		if _, exists := s.entries[token]; !exists {
			s.entries[token] = scanUploadTokenEntry{classID: classID, expiresAt: expiresAt}
			s.mu.Unlock()
			return token, expiresAt, nil
		}
		s.mu.Unlock()
	}
}

func (s *scanUploadTokenStore) consume(token string) (int, error) {
	token = strings.TrimSpace(token)
	if len(token) != 43 {
		return 0, ErrInvalidScanUploadToken
	}

	if global.GVA_CONFIG.System.UseRedis && global.GVA_REDIS != nil {
		classID, err := consumeScanUploadToken.Run(
			context.Background(),
			global.GVA_REDIS,
			[]string{scanUploadTokenPrefix + token},
		).Text()
		if errors.Is(err, redis.Nil) {
			return 0, ErrInvalidScanUploadToken
		}
		if err != nil {
			return 0, fmt.Errorf("读取扫码上传凭证失败: %w", err)
		}
		parsedClassID, err := strconv.Atoi(classID)
		if err != nil {
			return 0, ErrInvalidScanUploadToken
		}
		return parsedClassID, nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	entry, exists := s.entries[token]
	delete(s.entries, token)
	if !exists || !s.now().Before(entry.expiresAt) {
		return 0, ErrInvalidScanUploadToken
	}
	return entry.classID, nil
}

func (s *scanUploadTokenStore) removeExpiredLocked() {
	now := s.now()
	for token, entry := range s.entries {
		if !now.Before(entry.expiresAt) {
			delete(s.entries, token)
		}
	}
}

func (e *FileUploadAndDownloadService) IssueScanUploadToken(classID int) (string, time.Time, error) {
	return scanUploadTokens.issue(classID)
}

func (e *FileUploadAndDownloadService) ConsumeScanUploadToken(token string) (int, error) {
	return scanUploadTokens.consume(token)
}
