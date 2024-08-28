package biz_apphub

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/jsonx"
	"github.com/flipped-aurora/gin-vue-admin/server/service/biz_apphub"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

// GetDeployList 分页获取biz_apphub_record列表
// @Tags BizAppHub
// @Summary 分页获取biz_apphub_record列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizAppHubSearch true "分页获取biz_apphub列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bizAppHub/getDeployList [get]
func (bizAppHubApi *BizAppHubApi) GetDeployList(c *gin.Context) {
	var pageInfo biz_apphubReq.GetDeployList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.OperateUser = c.GetString("user")
	list, total, err := bizAppHubService.GetDeployList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// RollbackVersion 回滚版本biz_apphub
// @Tags BizAppHub
// @Summary 回滚版本biz_apphub
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizAppHub true "回滚版本biz_apphub"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bizAppHub/rollbackVersion [put]
func (bizAppHubApi *BizAppHubApi) RollbackVersion(c *gin.Context) {
	var bizAppHub biz_apphubReq.RollbackVersion
	err := c.ShouldBindJSON(&bizAppHub)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//bizAppHub.OperateUser=c.GetString("user")
	//bizAppHub.UpdatedBy = utils.GetUserID(c)
	bizAppHub.OperateUser = c.GetString("user")
	err = bizAppHubService.RollbackVersion(bizAppHub)
	if err != nil {
		global.GVA_LOG.Error("回滚版本失败!", zap.Error(err))
		response.FailWithMessage("回滚版本失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("回滚版本成功", c)
}

// PostCall 后端命令接口
// @Tags BizAppHub
// @Summary 后端命令接口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizAppHub true "后端命令接口"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bizAppHub/api/cmd/call/:user/:soft/:command [post]
func (bizAppHubApi *BizAppHubApi) PostCall(c *gin.Context) {
	var (
		req biz_apphubReq.Call
		err error
	)
	req.Soft = c.Param("soft")
	req.Command = c.Param("command")
	req.Method = "POST"
	req.User = c.Param("user")
	if req.Soft == "" {
		response.FailWithMessage("soft不能为空", c)
		return
	}
	if req.Command == "" {
		response.FailWithMessage("Command不能为空", c)
		return
	}

	if req.User == "" {
		response.FailWithMessage("请登录后访问", c)
		return
	}
	err = c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//bizAppHub.OperateUser=c.GetString("user")
	//bizAppHub.UpdatedBy = utils.GetUserID(c)
	//bizAppHub.OperateUser = c.GetString("user")
	//todo 验证用户是否有调用该应用的权限
	j, err := req.RequestJSON()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.ReqBody = j
	caller := biz_apphub.NewCaller("")
	req.RequestJsonPath = req.GetRequestFilePath(caller.CallerPath())
	if c.Request.Header.Get("content-type") == "multipart/form-data" { //如果上传的有文件，需要下载文件，然后copy到
		files, err := getFileMap(c, caller.CallerPath())
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		req.Files = files
		value := c.PostForm("jsonData")
		if value != "" {
			mp := make(map[string]interface{})
			err := json.Unmarshal([]byte(value), &mp)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
			req.Data = mp
		}
	}

	err = jsonx.SaveFile(req.RequestJsonPath, req) //
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	defer os.Remove(req.RequestJsonPath)

	call, err := caller.Call(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if call.StatusCode != 200 {
		c.JSON(call.StatusCode, gin.H{"msg": call.Data})
		return
	}
	if call.ContentType == "file" {
		if call.HasFile {
			//if call.HasFile
			fileName := filepath.Base(call.FilePath) // 获取文件名
			if call.DeleteFile {
				defer os.Remove(call.FilePath)
			}
			// 如果请求中有自定义文件名，则使用自定义文件名
			if customFileName := c.Query("filename"); customFileName != "" {
				fileName = customFileName
			}
			c.Writer.Header().Add("Content-Disposition", "attachment; filename="+fileName)
			c.File(call.FilePath)
			return
		}
	}

	if call.ContentType == "json" {
		c.Data(200, "application/json; charset=utf-8", []byte(jsonx.JSONString(call.Data)))
	}
	if call.ContentType == "text" {
		c.Data(200, "text/plain; charset=utf-8", []byte(fmt.Sprintf("%v", call.Data)))
	}
}

// GetCall 后端命令接口
// @Tags BizAppHub
// @Summary 后端命令接口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizAppHub true "后端命令接口"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bizAppHub/api/cmd/call/:user/:soft/:command [post]
func (bizAppHubApi *BizAppHubApi) GetCall(c *gin.Context) {
	var (
		req biz_apphubReq.Call
		err error
	)
	req.Soft = c.Param("soft")
	req.Command = c.Param("command")
	req.User = c.Param("user")
	req.Method = "GET"
	if req.Soft == "" {
		response.FailWithMessage("soft不能为空", c)
		return
	}
	if req.Command == "" {
		response.FailWithMessage("Command不能为空", c)
		return
	}

	if req.User == "" {
		response.FailWithMessage("user 不能为空", c)
		return
	}
	if req.Data == nil {
		req.Data = make(map[string]interface{})
	}
	queryMap := c.Request.URL.Query()
	for k, values := range queryMap {
		if len(values) > 1 {
			req.Data[k] = values
		} else {
			req.Data[k] = values[0]
		}
	}
	//todo 验证用户是否有调用该应用的权限
	j, err := req.RequestJSON()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.ReqBody = j
	caller := biz_apphub.NewCaller("")
	req.RequestJsonPath = req.GetRequestFilePath(caller.CallerPath())

	err = jsonx.SaveFile(req.RequestJsonPath, req) //
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	defer os.Remove(req.RequestJsonPath)

	call, err := caller.Call(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if call.StatusCode != 200 {
		c.JSON(call.StatusCode, gin.H{
			"msg": call.Data,
		})
		return
	}

	if call.ContentType == "file" {
		if call.HasFile {
			//if call.HasFile
			fileName := filepath.Base(call.FilePath) // 获取文件名
			if call.DeleteFile {
				defer os.Remove(call.FilePath)
			}
			// 如果请求中有自定义文件名，则使用自定义文件名
			if customFileName := c.Query("filename"); customFileName != "" {
				fileName = customFileName
			}
			c.Writer.Header().Add("Content-Disposition", "attachment; filename="+fileName)
			c.File(call.FilePath)
			return
		}
	}

	if call.ContentType == "json" {
		c.Data(200, "application/json; charset=utf-8", []byte(jsonx.JSONString(call.Data)))
	}
	if call.ContentType == "text" {
		c.Data(200, "text/plain; charset=utf-8", []byte(fmt.Sprintf("%v", call.Data)))
	}
}

func getFileMap(c *gin.Context, fileRoot string) (filePath []string, err error) {
	err = c.Request.ParseMultipartForm(32 << 20) // 32MB
	if err != nil {
		//http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	files := c.Request.MultipartForm.File["file"]
	for _, fileHeader := range files {
		filePathEl := fmt.Sprintf("%v\\file\\%v", fileRoot, time.Now().UnixNano())
		filePathAbs := filePathEl + "\\" + fileHeader.Filename
		os.MkdirAll(filePathEl, os.ModePerm)
		filePath = append(filePath, filePathAbs)
		err := c.SaveUploadedFile(fileHeader, filePathAbs)
		if err != nil {
			return nil, err
		}
	}
	return filePath, nil
}
