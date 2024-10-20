package biz_apphub

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	biz_apphubReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
)

type BizCloudFunctionApi struct{}

type Param struct {
	Code     string `json:"code"`
	Desc     string `json:"desc"`
	Mode     string `json:"mode"`
	Type     string `json:"type"`
	Required string `json:"required,omitempty"`
	MockData string `json:"mock_data"`
}

// CreateBizCloudFunction 创建云函数
// @Tags BizCloudFunction
// @Summary 创建云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizCloudFunction true "创建云函数"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /bizCloudFunction/createBizCloudFunction [post]
func (bizCloudFunctionApi *BizCloudFunctionApi) CreateBizCloudFunction(c *gin.Context) {
	var bizCloudFunction biz_apphub.BizCloudFunction
	var runner biz_apphub.BizToolCmdSrvApi
	err := c.ShouldBindJSON(&bizCloudFunction)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bizCloudFunction.TenantUser = c.GetString("user")
	//mp:=make(map[string]string)
	type ApiConfig struct {
		Path   string `json:"path"` //api/runner/run/beiluo/apphub/array/SplitJoin
		Method string `json:"method"`
	}
	a := ApiConfig{}
	var params []Param
	err = json.Unmarshal(bizCloudFunction.ApiConfig, &a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = json.Unmarshal(bizCloudFunction.Param, &params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	infiles := []string{}
	outfiles := []string{}
	for _, param := range params {
		if param.Type == "file" {
			if strings.ToLower(param.Mode) == "in" {
				infiles = append(infiles, param.Code)
			} else if strings.ToLower(param.Mode) == "out" {
				outfiles = append(outfiles, param.Code)
			}
		}
	}
	if len(infiles) != 0 {
		bizCloudFunction.InFile = strings.Join(infiles, ";")
	}
	if len(outfiles) != 0 {
		bizCloudFunction.OutFile = strings.Join(outfiles, ";")
	}

	a.Path = "api/runner/run/" + bizCloudFunction.TenantUser + "/" + a.Path
	marshal, err := json.Marshal(a)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bizCloudFunction.ApiConfig = marshal
	//bizCloudFunction.ApiConfig=
	config := bizCloudFunction.GetApiConfig()
	info, err := config.GetRunnerInfo()
	if err != nil {
		err = nil
	} else {
		err := global.GVA_DB.Model(&biz_apphub.BizToolCmdSrvApi{}).
			Where("tenant_user =? AND app_code=?", info.User, info.Code).First(&runner).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			err = nil
		} else {
			bizCloudFunction.RunnerID = runner.ID
			bizCloudFunction.ApiPath = info.Path
			bizCloudFunction.ApiMethod = config.Method
			bizCloudFunction.ApiFullPath = config.Path
		}
	}

	err = bizCloudFunctionService.CreateBizCloudFunction(&bizCloudFunction)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBizCloudFunction 删除云函数
// @Tags BizCloudFunction
// @Summary 删除云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizCloudFunction true "删除云函数"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /bizCloudFunction/deleteBizCloudFunction [delete]
func (bizCloudFunctionApi *BizCloudFunctionApi) DeleteBizCloudFunction(c *gin.Context) {
	ID := c.Query("ID")
	err := bizCloudFunctionService.DeleteBizCloudFunction(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBizCloudFunctionByIds 批量删除云函数
// @Tags BizCloudFunction
// @Summary 批量删除云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /bizCloudFunction/deleteBizCloudFunctionByIds [delete]
func (bizCloudFunctionApi *BizCloudFunctionApi) DeleteBizCloudFunctionByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := bizCloudFunctionService.DeleteBizCloudFunctionByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBizCloudFunction 更新云函数
// @Tags BizCloudFunction
// @Summary 更新云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz_apphub.BizCloudFunction true "更新云函数"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /bizCloudFunction/updateBizCloudFunction [put]
func (bizCloudFunctionApi *BizCloudFunctionApi) UpdateBizCloudFunction(c *gin.Context) {
	var bizCloudFunction biz_apphub.BizCloudFunction
	var runner biz_apphub.BizToolCmdSrvApi

	err := c.ShouldBindJSON(&bizCloudFunction)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	bizCloudFunction.InFile = ""
	bizCloudFunction.OutFile = ""

	config := bizCloudFunction.GetApiConfig()
	info, err := config.GetRunnerInfo()
	if err != nil {
		err = nil
	} else {
		err := global.GVA_DB.Model(&biz_apphub.BizToolCmdSrvApi{}).
			Where("tenant_user =? AND app_code=?", info.User, info.Code).First(&runner).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			err = nil
		} else {
			bizCloudFunction.RunnerID = runner.ID
			bizCloudFunction.ApiPath = info.Path
			bizCloudFunction.ApiMethod = config.Method
			bizCloudFunction.ApiFullPath = config.Path
		}
	}

	var params []Param
	err = json.Unmarshal(bizCloudFunction.Param, &params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	infiles := []string{}
	outfiles := []string{}
	for _, param := range params {
		if param.Type == "file" {
			if strings.ToLower(param.Mode) == "in" {
				infiles = append(infiles, param.Code)
			} else if strings.ToLower(param.Mode) == "out" {
				outfiles = append(outfiles, param.Code)
			}
		}
	}
	if len(infiles) != 0 {
		bizCloudFunction.InFile = strings.Join(infiles, ";")
	}
	if len(outfiles) != 0 {
		bizCloudFunction.OutFile = strings.Join(outfiles, ";")
	}

	err = bizCloudFunctionService.UpdateBizCloudFunction(bizCloudFunction)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBizCloudFunction 用id查询云函数
// @Tags BizCloudFunction
// @Summary 用id查询云函数
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphub.BizCloudFunction true "用id查询云函数"
// @Success 200 {object} response.Response{data=biz_apphub.BizCloudFunction,msg=string} "查询成功"
// @Router /bizCloudFunction/findBizCloudFunction [get]
func (bizCloudFunctionApi *BizCloudFunctionApi) FindBizCloudFunction(c *gin.Context) {
	ID := c.Query("ID")
	rebizCloudFunction, err := bizCloudFunctionService.GetBizCloudFunction(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rebizCloudFunction, c)
}

// GetBizCloudFunctionList 分页获取云函数列表
// @Tags BizCloudFunction
// @Summary 分页获取云函数列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizCloudFunctionSearch true "分页获取云函数列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /bizCloudFunction/getBizCloudFunctionList [get]
func (bizCloudFunctionApi *BizCloudFunctionApi) GetBizCloudFunctionList(c *gin.Context) {
	var pageInfo biz_apphubReq.BizCloudFunctionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := bizCloudFunctionService.GetBizCloudFunctionInfoList(pageInfo)
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

// GetBizCloudFunctionPublic 不需要鉴权的云函数接口
// @Tags BizCloudFunction
// @Summary 不需要鉴权的云函数接口
// @accept application/json
// @Produce application/json
// @Param data query biz_apphubReq.BizCloudFunctionSearch true "分页获取云函数列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /bizCloudFunction/getBizCloudFunctionPublic [get]
func (bizCloudFunctionApi *BizCloudFunctionApi) GetBizCloudFunctionPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的云函数接口信息",
	}, "获取成功", c)
}

func (bizCloudFunctionApi *BizCloudFunctionApi) SyncFunction(c *gin.Context) {
	req := global.GVA_MODEL{}

	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	function, err := bizCloudFunctionService.SyncFunction(&biz_apphub.BizToolCmdSrvApi{
		GVA_MODEL: global.GVA_MODEL{ID: req.ID},
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(function, c)
	//bizCloudFunctionService.GetBizCloudFunction()
}
