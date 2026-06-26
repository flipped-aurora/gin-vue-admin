package media

import (
	"crypto/md5"
	"encoding/hex"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media/request"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func testDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&media.MediaUpload{}, &media.MediaUploadChunk{}, &media.FileUploadAndDownload{}); err != nil {
		t.Fatal(err)
	}
	global.GVA_DB = db
}

func TestInitNewAndResume(t *testing.T) {
	testDB(t)
	s := &MediaUploadService{}
	req := request.UploadInitReq{FileName: "a.bin", FileHash: "h1", FileSize: 10, ChunkSize: 5, ChunkTotal: 2}
	r1, err := s.Init(1, req)
	if err != nil || r1.Instant || r1.UploadID == 0 {
		t.Fatalf("init new: %+v err=%v", r1, err)
	}
	// 模拟收到分片 0
	global.GVA_DB.Create(&media.MediaUploadChunk{UploadID: r1.UploadID, ChunkIndex: 0})
	r2, _ := s.Init(1, req)
	if r2.UploadID != r1.UploadID || len(r2.UploadedChunks) != 1 || r2.UploadedChunks[0] != 0 {
		t.Fatalf("resume should report chunk 0: %+v", r2)
	}
}

func TestSaveChunkIdempotent(t *testing.T) {
	testDB(t)
	global.GVA_CONFIG.Media.ChunkDir = t.TempDir()
	s := &MediaUploadService{}
	r, _ := s.Init(1, request.UploadInitReq{FileName: "a", FileHash: "h", FileSize: 3, ChunkSize: 3, ChunkTotal: 1})
	data := []byte("abc")
	import_md5 := "900150983cd24fb0d6963f7d28e17f72" // md5("abc")
	if err := s.SaveChunk(1, r.UploadID, 0, import_md5, data); err != nil {
		t.Fatal(err)
	}
	if err := s.SaveChunk(1, r.UploadID, 0, import_md5, data); err != nil {
		t.Fatalf("重复收片应幂等: %v", err)
	}
	var cnt int64
	global.GVA_DB.Model(&media.MediaUploadChunk{}).Where("upload_id = ?", r.UploadID).Count(&cnt)
	if cnt != 1 {
		t.Fatalf("应只有 1 条分片记录, got %d", cnt)
	}
	// 错误 hash 应失败
	if err := s.SaveChunk(1, r.UploadID, 0, "deadbeef", data); err == nil {
		t.Fatal("hash 不符应报错")
	}
}

func TestCompleteMergeAndRegister(t *testing.T) {
	testDB(t)
	global.GVA_CONFIG.Media.ChunkDir = t.TempDir()
	global.GVA_CONFIG.System.OssType = "local"
	global.GVA_CONFIG.Local.StorePath = t.TempDir()
	global.GVA_CONFIG.Local.Path = "uploads/file"
	s := &MediaUploadService{}
	full := []byte("hello big world")
	// 整文件 md5
	r, _ := s.Init(1, request.UploadInitReq{FileName: "x.txt", FileHash: md5hex(full), FileSize: int64(len(full)), ChunkSize: 6, ChunkTotal: 3})
	chunks := [][]byte{full[0:6], full[6:10], full[10:]}
	for i, c := range chunks {
		if err := s.SaveChunk(1, r.UploadID, i, md5hex(c), c); err != nil {
			t.Fatal(err)
		}
	}
	m, err := s.Complete(1, r.UploadID)
	if err != nil {
		t.Fatalf("complete: %v", err)
	}
	if m.ID == 0 || m.Url == "" {
		t.Fatalf("media not registered: %+v", m)
	}
	var up media.MediaUpload
	global.GVA_DB.First(&up, r.UploadID)
	if up.Status != media.UploadStatusCompleted || up.MediaID != m.ID {
		t.Fatalf("session not completed: %+v", up)
	}
}

func md5hex(b []byte) string {
	sum := md5.Sum(b)
	return hex.EncodeToString(sum[:])
}
