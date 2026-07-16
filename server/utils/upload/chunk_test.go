package upload

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func setupCfg(t *testing.T) {
	dir := t.TempDir()
	global.GVA_CONFIG.Media.ChunkDir = dir
}

func TestSaveAndMergeChunks(t *testing.T) {
	setupCfg(t)
	parts := [][]byte{[]byte("hello "), []byte("big "), []byte("world")}
	for i, p := range parts {
		if _, err := SaveChunkFile(7, i, p); err != nil {
			t.Fatalf("save %d: %v", i, err)
		}
	}
	dst := filepath.Join(t.TempDir(), "merged.bin")
	got, err := MergeChunks(7, len(parts), dst)
	if err != nil {
		t.Fatalf("merge: %v", err)
	}
	data, _ := os.ReadFile(dst)
	if string(data) != "hello big world" {
		t.Fatalf("merged content = %q", data)
	}
	sum := md5.Sum([]byte("hello big world"))
	if got != hex.EncodeToString(sum[:]) {
		t.Fatalf("md5 mismatch: %s", got)
	}
	if err := RemoveUploadDir(7); err != nil {
		t.Fatalf("remove: %v", err)
	}
	if _, err := os.Stat(ChunkDir(7)); !os.IsNotExist(err) {
		t.Fatalf("dir should be gone")
	}
	_ = config.Media{} // ensure import used
}
