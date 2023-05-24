package tcb

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"
)

const (
	// 获取文件上传链接
	uploadFilePathURL = "https://api.weixin.qq.com/tcb/uploadfile"
	// 获取文件下载链接
	batchDownloadFileURL = "https://api.weixin.qq.com/tcb/batchdownloadfile"
	// 删除文件链接
	batchDeleteFileURL = "https://api.weixin.qq.com/tcb/batchdeletefile"
)

// UploadFileReq 上传文件请求值
type UploadFileReq struct {
	Env  string `json:"env,omitempty"`
	Path string `json:"path,omitempty"`
}

// UploadFileRes 上传文件返回结果
type UploadFileRes struct {
	util.CommonError
	URL           string `json:"url"`           // 上传url
	Token         string `json:"token"`         // token
	Authorization string `json:"authorization"` // authorization
	FileID        string `json:"file_id"`       // 文件ID
	CosFileID     string `json:"cos_file_id"`   // cos文件ID
}

// BatchDownloadFileReq 上传文件请求值
type BatchDownloadFileReq struct {
	Env      string          `json:"env,omitempty"`
	FileList []*DownloadFile `json:"file_list,omitempty"`
}

// DownloadFile 文件信息
type DownloadFile struct {
	FileID string `json:"fileid"`  // 文件ID
	MaxAge int64  `json:"max_age"` // 下载链接有效期
}

// BatchDownloadFileRes 上传文件返回结果
type BatchDownloadFileRes struct {
	util.CommonError
	FileList []struct {
		FileID      string `json:"file_id"`      // 文件ID
		DownloadURL string `json:"download_url"` // 下载链接
		Status      int64  `json:"status"`       // 状态码
		ErrMsg      string `json:"errmsg"`       // 该文件错误信息
	} `json:"file_list"`
}

// BatchDeleteFileReq 批量删除文件请求参数
type BatchDeleteFileReq struct {
	Env        string   `json:"env,omitempty"`
	FileIDList []string `json:"fileid_list,omitempty"`
}

// BatchDeleteFileRes 批量删除文件返回结果
type BatchDeleteFileRes struct {
	util.CommonError
	DeleteList []struct {
		FileID string `json:"fileid"`
		Status int64  `json:"status"`
		ErrMsg string `json:"errmsg"`
	} `json:"delete_list"`
}

// UploadFile 上传文件
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/storage/uploadFile.html
func (tcb *Tcb) UploadFile(env, path string) (*UploadFileRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", uploadFilePathURL, accessToken)
	req := &UploadFileReq{
		Env:  env,
		Path: path,
	}
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}
	uploadFileRes := &UploadFileRes{}
	err = util.DecodeWithError(response, uploadFileRes, "UploadFile")
	return uploadFileRes, err
}

// BatchDownloadFile 获取文件下载链接
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/storage/batchDownloadFile.html
func (tcb *Tcb) BatchDownloadFile(env string, fileList []*DownloadFile) (*BatchDownloadFileRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", batchDownloadFileURL, accessToken)
	req := &BatchDownloadFileReq{
		Env:      env,
		FileList: fileList,
	}
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}
	batchDownloadFileRes := &BatchDownloadFileRes{}
	err = util.DecodeWithError(response, batchDownloadFileRes, "BatchDownloadFile")
	return batchDownloadFileRes, err
}

// BatchDeleteFile 批量删除文件
//
//reference:https://developers.weixin.qq.com/miniprogram/dev/wxcloud/reference-http-api/storage/batchDeleteFile.html
func (tcb *Tcb) BatchDeleteFile(env string, fileIDList []string) (*BatchDeleteFileRes, error) {
	accessToken, err := tcb.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", batchDeleteFileURL, accessToken)
	req := &BatchDeleteFileReq{
		Env:        env,
		FileIDList: fileIDList,
	}
	response, err := util.PostJSON(uri, req)
	if err != nil {
		return nil, err
	}
	batchDeleteFileRes := &BatchDeleteFileRes{}
	err = util.DecodeWithError(response, batchDeleteFileRes, "BatchDeleteFile")
	return batchDeleteFileRes, err
}
