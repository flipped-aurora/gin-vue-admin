package store

// FileStore 存储接口
type FileStore interface {
	// FileSave 保存文件
	FileSave(localFilePath string, savePath string) (*FileSaveInfo, error)

	// GetFile 下载文件
	GetFile(savePath string) (*GetFileInfo, error)

	// DeleteFile 删除文件
	DeleteFile(savePath string) error
}

type FileSaveInfo struct {
	SavePath     string `json:"save_path"`      //   /beiluo/d.zip
	SaveFullPath string `json:"save_full_path"` //http://cdn.geeleo.com/beiluo/d.zip
	FileName     string `json:"file_name"`
	FileType     string `json:"file_type"`
}

type GetFileInfo struct {
	FileSaveInfo
	FileLocalPath string `json:"file_local_path"`
	FileSize      int64  `json:"file_size"`
	FileType      string `json:"file_type"`
}
