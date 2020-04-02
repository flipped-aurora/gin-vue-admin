package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/dbModel"
	"gin-vue-admin/model/modelInterface"
	"github.com/gin-gonic/gin"
	"strings"
)

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/upload [post]
func UploadFile(c *gin.Context) {
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("上传文件失败，%v", err), gin.H{})
	} else {
		//文件上传后拿到文件路径
		err, filePath, key := servers.Upload(header, USER_HEADER_BUCKET, USER_HEADER_IMG_PATH)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("接收返回值失败，%v", err), gin.H{})
		} else {
			//修改数据库后得到修改后的user并且返回供前端使用
			var file dbModel.ExaFileUploadAndDownload
			file.Url = filePath
			file.Name = header.Filename
			s := strings.Split(file.Name, ".")
			file.Tag = s[len(s)-1]
			file.Key = key
			if noSave == "0" {
				err = file.Upload()
			}
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("修改数据库链接失败，%v", err), gin.H{})
			} else {
				servers.ReportFormat(c, true, "上传成功", gin.H{"file": file})
			}
		}
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body dbModel.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /fileUploadAndDownload/deleteFile [post]
func DeleteFile(c *gin.Context) {
	var file dbModel.ExaFileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	err, f := file.FindFile()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("删除失败，%v", err), gin.H{})
	} else {
		err = servers.DeleteFile(USER_HEADER_BUCKET, f.Key)
		if err != nil {
			servers.ReportFormat(c, false, fmt.Sprintf("删除失败，%v", err), gin.H{})
		} else {
			err = f.DeleteFile()
			if err != nil {
				servers.ReportFormat(c, false, fmt.Sprintf("删除失败，%v", err), gin.H{})
			} else {
				servers.ReportFormat(c, true, fmt.Sprintf("删除成功，%v", err), gin.H{})
			}
		}
	}
}

// @Tags ExaFileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取文件户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fileUploadAndDownload/getFileList [post]
func GetFileList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(dbModel.ExaFileUploadAndDownload).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"list":     list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}
