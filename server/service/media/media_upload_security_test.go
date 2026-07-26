package media

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/media"
	"github.com/flipped-aurora/gin-vue-admin/server/model/media/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
)

func TestUploadFileRejectsActiveContentBeforeStorage(t *testing.T) {
	oldSystem := global.GVA_CONFIG.System
	oldLocal := global.GVA_CONFIG.Local
	global.GVA_CONFIG.System.OssType = "local"
	global.GVA_CONFIG.Local.StorePath = t.TempDir()
	global.GVA_CONFIG.Local.Path = "/uploads/file"
	t.Cleanup(func() {
		global.GVA_CONFIG.System = oldSystem
		global.GVA_CONFIG.Local = oldLocal
	})

	source := filepath.Join(t.TempDir(), "source")
	if err := os.WriteFile(source, []byte("<script>alert(1)</script>"), 0600); err != nil {
		t.Fatal(err)
	}
	header, cleanup, err := upload.BuildFileHeader(source, "file", "payload.html")
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	if _, err = FileUploadAndDownloadServiceApp.UploadFile(context.Background(), header, "1", 0, 1); err == nil {
		t.Fatal("UploadFile() error = nil, want active-content rejection")
	}
	entries, err := os.ReadDir(global.GVA_CONFIG.Local.StorePath)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Fatalf("UploadFile() wrote %d files before rejecting active content", len(entries))
	}
}

func TestInitRejectsActiveContentBeforeSessionCreation(t *testing.T) {
	testDB(t)
	s := &MediaUploadService{}
	_, err := s.Init(context.Background(), 1, request.UploadInitReq{
		FileName: "payload.svg", FileHash: "hash", FileSize: 10, ChunkSize: 10, ChunkTotal: 1,
	})
	if err == nil {
		t.Fatal("Init() error = nil, want active-content rejection")
	}
	var count int64
	if err := global.GVA_DB.Model(&model.MediaUpload{}).Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatalf("Init() created %d upload sessions for rejected content", count)
	}
}

func TestCompleteRejectsActiveContentBeforeStateChange(t *testing.T) {
	testDB(t)
	s := &MediaUploadService{}
	up := model.MediaUpload{
		UserID: 1, FileName: "payload.html", FileHash: "hash", FileSize: 10,
		ChunkSize: 10, ChunkTotal: 1, Status: model.UploadStatusUploading,
	}
	if err := global.GVA_DB.Create(&up).Error; err != nil {
		t.Fatal(err)
	}
	if _, err := s.Complete(context.Background(), 1, up.ID); err == nil {
		t.Fatal("Complete() error = nil, want active-content rejection")
	}
	if err := global.GVA_DB.First(&up, up.ID).Error; err != nil {
		t.Fatal(err)
	}
	if up.Status != model.UploadStatusUploading {
		t.Fatalf("Complete() status = %q, want %q", up.Status, model.UploadStatusUploading)
	}
}
