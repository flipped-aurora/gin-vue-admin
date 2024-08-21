package oss

type Store interface {
	UploadLocalFile(localFilePath string, ossPath string) (*File, error)
	GetUploadToken() string
}
