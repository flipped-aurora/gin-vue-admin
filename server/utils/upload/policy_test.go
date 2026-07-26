package upload

import "testing"

func TestValidateFileExtension(t *testing.T) {
	allowed := []string{
		"photo.jpg", "photo.JPEG", "photo.png", "photo.webp", "photo.avif",
		"audio.mp3", "audio.flac", "video.mp4", "video.webm",
		"notes.txt", "data.csv", "data.json", "manual.pdf",
		"document.docx", "sheet.xlsx", "slides.pptx", "archive.zip", "blob.bin",
	}
	for _, name := range allowed {
		t.Run("allow_"+name, func(t *testing.T) {
			if err := ValidateFileExtension(name); err != nil {
				t.Fatalf("ValidateFileExtension(%q) error = %v", name, err)
			}
		})
	}

	rejected := []string{
		"no-extension", "page.html", "page.HTML", "page.xhtml", "vector.svg",
		"document.xml", "script.js", "module.mjs", "style.css", "photo.png.html",
		"payload.exe", "trailing.jpg ", "../page.html",
	}
	for _, name := range rejected {
		t.Run("reject_"+name, func(t *testing.T) {
			if err := ValidateFileExtension(name); err == nil {
				t.Fatalf("ValidateFileExtension(%q) error = nil, want rejection", name)
			}
		})
	}
}

func TestCanServeInline(t *testing.T) {
	for _, name := range []string{"photo.png", "photo.JPEG", "audio.mp3", "video.mp4"} {
		if !CanServeInline(name) {
			t.Errorf("CanServeInline(%q) = false, want true", name)
		}
	}
	for _, name := range []string{"manual.pdf", "notes.txt", "archive.zip", "legacy.html", "vector.svg"} {
		if CanServeInline(name) {
			t.Errorf("CanServeInline(%q) = true, want false", name)
		}
	}
}
