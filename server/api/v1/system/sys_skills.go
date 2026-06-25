package system

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SkillsApi struct{}

// GetTools
// @Tags      Skills
// @Summary   获取工具列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=object,msg=string}  "获取工具列表"
// @Router    /skills/getTools [get]
func (s *SkillsApi) GetTools(c *gin.Context) {
	data, err := skillsService.Tools(c.Request.Context())
	if err != nil {
		global.GVA_LOG.Error("获取工具列表失败", zap.Error(err))
		response.FailWithMessage("获取工具列表失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"tools": data}, "获取成功", c)
}

// GetSkillList
// @Tags      Skills
// @Summary   获取技能列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillToolRequest                    true  "工具标识"
// @Success   200   {object}  response.Response{data=object,msg=string}   "获取技能列表"
// @Router    /skills/getSkillList [post]
func (s *SkillsApi) GetSkillList(c *gin.Context) {
	var req request.SkillToolRequest
	_ = c.ShouldBindJSON(&req)
	data, err := skillsService.List(c.Request.Context(), req.Tool)
	if err != nil {
		global.GVA_LOG.Error("获取技能列表失败", zap.Error(err))
		response.FailWithMessage("获取技能列表失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"skills": data}, "获取成功", c)
}

// GetSkillDetail
// @Tags      Skills
// @Summary   获取技能详情
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillDetailRequest                  true  "工具标识, 技能标识"
// @Success   200   {object}  response.Response{data=object,msg=string}   "获取技能详情"
// @Router    /skills/getSkillDetail [post]
func (s *SkillsApi) GetSkillDetail(c *gin.Context) {
	var req request.SkillDetailRequest
	_ = c.ShouldBindJSON(&req)
	data, err := skillsService.Detail(c.Request.Context(), req.Tool, req.Skill)
	if err != nil {
		global.GVA_LOG.Error("获取技能详情失败", zap.Error(err))
		response.FailWithMessage("获取技能详情失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"detail": data}, "获取成功", c)
}

// SaveSkill
// @Tags      Skills
// @Summary   保存技能
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillSaveRequest       true  "工具标识, 技能标识, 元数据, 内容"
// @Success   200   {object}  response.Response{msg=string}  "保存技能"
// @Router    /skills/saveSkill [post]
func (s *SkillsApi) SaveSkill(c *gin.Context) {
	var req request.SkillSaveRequest
	_ = c.ShouldBindJSON(&req)
	if err := skillsService.Save(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("保存技能失败", zap.Error(err))
		response.FailWithMessage("保存技能失败", c)
		return
	}
	response.OkWithMessage("保存成功", c)
}

// DeleteSkill
// @Tags      Skills
// @Summary   删除技能
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillDeleteRequest     true  "工具标识, 技能标识"
// @Success   200   {object}  response.Response{msg=string}  "删除技能"
// @Router    /skills/deleteSkill [post]
func (s *SkillsApi) DeleteSkill(c *gin.Context) {
	var req request.SkillDeleteRequest
	_ = c.ShouldBindJSON(&req)
	if err := skillsService.Delete(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("删除技能失败", zap.Error(err))
		response.FailWithMessage("删除技能失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// CreateScript
// @Tags      Skills
// @Summary   创建脚本
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillScriptCreateRequest            true  "工具标识, 技能标识, 文件名, 脚本类型"
// @Success   200   {object}  response.Response{data=object,msg=string}   "创建脚本"
// @Router    /skills/createScript [post]
func (s *SkillsApi) CreateScript(c *gin.Context) {
	var req request.SkillScriptCreateRequest
	_ = c.ShouldBindJSON(&req)
	fileName, content, err := skillsService.CreateScript(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("创建脚本失败", zap.Error(err))
		response.FailWithMessage("创建脚本失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"fileName": fileName, "content": content}, "创建成功", c)
}

// GetScript
// @Tags      Skills
// @Summary   读取脚本
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillFileRequest                    true  "工具标识, 技能标识, 文件名"
// @Success   200   {object}  response.Response{data=object,msg=string}   "读取脚本"
// @Router    /skills/getScript [post]
func (s *SkillsApi) GetScript(c *gin.Context) {
	var req request.SkillFileRequest
	_ = c.ShouldBindJSON(&req)
	content, err := skillsService.GetScript(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("读取脚本失败", zap.Error(err))
		response.FailWithMessage("读取脚本失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"content": content}, "获取成功", c)
}

// SaveScript
// @Tags      Skills
// @Summary   保存脚本
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillFileSaveRequest   true  "工具标识, 技能标识, 文件名, 内容"
// @Success   200   {object}  response.Response{msg=string}  "保存脚本"
// @Router    /skills/saveScript [post]
func (s *SkillsApi) SaveScript(c *gin.Context) {
	var req request.SkillFileSaveRequest
	_ = c.ShouldBindJSON(&req)
	if err := skillsService.SaveScript(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("保存脚本失败", zap.Error(err))
		response.FailWithMessage("保存脚本失败", c)
		return
	}
	response.OkWithMessage("保存成功", c)
}

// CreateResource
// @Tags      Skills
// @Summary   创建资源
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillResourceCreateRequest          true  "工具标识, 技能标识, 文件名"
// @Success   200   {object}  response.Response{data=object,msg=string}   "创建资源"
// @Router    /skills/createResource [post]
func (s *SkillsApi) CreateResource(c *gin.Context) {
	var req request.SkillResourceCreateRequest
	_ = c.ShouldBindJSON(&req)
	fileName, content, err := skillsService.CreateResource(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("创建资源失败", zap.Error(err))
		response.FailWithMessage("创建资源失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"fileName": fileName, "content": content}, "创建成功", c)
}

// GetResource
// @Tags      Skills
// @Summary   读取资源
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillFileRequest                    true  "工具标识, 技能标识, 文件名"
// @Success   200   {object}  response.Response{data=object,msg=string}   "读取资源"
// @Router    /skills/getResource [post]
func (s *SkillsApi) GetResource(c *gin.Context) {
	var req request.SkillFileRequest
	_ = c.ShouldBindJSON(&req)
	content, err := skillsService.GetResource(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("读取资源失败", zap.Error(err))
		response.FailWithMessage("读取资源失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"content": content}, "获取成功", c)
}

// SaveResource
// @Tags      Skills
// @Summary   保存资源
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillFileSaveRequest   true  "工具标识, 技能标识, 文件名, 内容"
// @Success   200   {object}  response.Response{msg=string}  "保存资源"
// @Router    /skills/saveResource [post]
func (s *SkillsApi) SaveResource(c *gin.Context) {
	var req request.SkillFileSaveRequest
	_ = c.ShouldBindJSON(&req)
	if err := skillsService.SaveResource(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("保存资源失败", zap.Error(err))
		response.FailWithMessage("保存资源失败", c)
		return
	}
	response.OkWithMessage("保存成功", c)
}

// CreateReference
// @Tags      Skills
// @Summary   创建参考
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillReferenceCreateRequest         true  "工具标识, 技能标识, 文件名"
// @Success   200   {object}  response.Response{data=object,msg=string}   "创建参考"
// @Router    /skills/createReference [post]
func (s *SkillsApi) CreateReference(c *gin.Context) {
	var req request.SkillReferenceCreateRequest
	_ = c.ShouldBindJSON(&req)
	fileName, content, err := skillsService.CreateReference(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("创建参考失败", zap.Error(err))
		response.FailWithMessage("创建参考失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"fileName": fileName, "content": content}, "创建成功", c)
}

// GetReference
// @Tags      Skills
// @Summary   读取参考
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillFileRequest                    true  "工具标识, 技能标识, 文件名"
// @Success   200   {object}  response.Response{data=object,msg=string}   "读取参考"
// @Router    /skills/getReference [post]
func (s *SkillsApi) GetReference(c *gin.Context) {
	var req request.SkillFileRequest
	_ = c.ShouldBindJSON(&req)
	content, err := skillsService.GetReference(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("读取参考失败", zap.Error(err))
		response.FailWithMessage("读取参考失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"content": content}, "获取成功", c)
}

// SaveReference
// @Tags      Skills
// @Summary   保存参考
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillFileSaveRequest   true  "工具标识, 技能标识, 文件名, 内容"
// @Success   200   {object}  response.Response{msg=string}  "保存参考"
// @Router    /skills/saveReference [post]
func (s *SkillsApi) SaveReference(c *gin.Context) {
	var req request.SkillFileSaveRequest
	_ = c.ShouldBindJSON(&req)
	if err := skillsService.SaveReference(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("保存参考失败", zap.Error(err))
		response.FailWithMessage("保存参考失败", c)
		return
	}
	response.OkWithMessage("保存成功", c)
}

// CreateTemplate
// @Tags      Skills
// @Summary   创建模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillTemplateCreateRequest          true  "工具标识, 技能标识, 文件名"
// @Success   200   {object}  response.Response{data=object,msg=string}   "创建模板"
// @Router    /skills/createTemplate [post]
func (s *SkillsApi) CreateTemplate(c *gin.Context) {
	var req request.SkillTemplateCreateRequest
	_ = c.ShouldBindJSON(&req)
	fileName, content, err := skillsService.CreateTemplate(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("创建模板失败", zap.Error(err))
		response.FailWithMessage("创建模板失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"fileName": fileName, "content": content}, "创建成功", c)
}

// GetTemplate
// @Tags      Skills
// @Summary   读取模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillFileRequest                    true  "工具标识, 技能标识, 文件名"
// @Success   200   {object}  response.Response{data=object,msg=string}   "读取模板"
// @Router    /skills/getTemplate [post]
func (s *SkillsApi) GetTemplate(c *gin.Context) {
	var req request.SkillFileRequest
	_ = c.ShouldBindJSON(&req)
	content, err := skillsService.GetTemplate(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("读取模板失败", zap.Error(err))
		response.FailWithMessage("读取模板失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"content": content}, "获取成功", c)
}

// SaveTemplate
// @Tags      Skills
// @Summary   保存模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillFileSaveRequest   true  "工具标识, 技能标识, 文件名, 内容"
// @Success   200   {object}  response.Response{msg=string}  "保存模板"
// @Router    /skills/saveTemplate [post]
func (s *SkillsApi) SaveTemplate(c *gin.Context) {
	var req request.SkillFileSaveRequest
	_ = c.ShouldBindJSON(&req)
	if err := skillsService.SaveTemplate(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("保存模板失败", zap.Error(err))
		response.FailWithMessage("保存模板失败", c)
		return
	}
	response.OkWithMessage("保存成功", c)
}

// GetGlobalConstraint
// @Tags      Skills
// @Summary   读取全局约束
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillToolRequest                    true  "工具标识"
// @Success   200   {object}  response.Response{data=object,msg=string}   "读取全局约束"
// @Router    /skills/getGlobalConstraint [post]
func (s *SkillsApi) GetGlobalConstraint(c *gin.Context) {
	var req request.SkillToolRequest
	_ = c.ShouldBindJSON(&req)
	content, exists, err := skillsService.GetGlobalConstraint(c.Request.Context(), req.Tool)
	if err != nil {
		global.GVA_LOG.Error("读取全局约束失败", zap.Error(err))
		response.FailWithMessage("读取全局约束失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"content": content, "exists": exists}, "获取成功", c)
}

// SaveGlobalConstraint
// @Tags      Skills
// @Summary   保存全局约束
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.SkillGlobalConstraintSaveRequest  true  "工具标识, 内容, 同步工具"
// @Success   200   {object}  response.Response{msg=string}             "保存全局约束"
// @Router    /skills/saveGlobalConstraint [post]
func (s *SkillsApi) SaveGlobalConstraint(c *gin.Context) {
	var req request.SkillGlobalConstraintSaveRequest
	_ = c.ShouldBindJSON(&req)
	if err := skillsService.SaveGlobalConstraint(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("保存全局约束失败", zap.Error(err))
		response.FailWithMessage("保存全局约束失败", c)
		return
	}
	response.OkWithMessage("保存成功", c)
}

// PackageSkill
// @Tags      Skills
// @Summary   打包技能
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/octet-stream
// @Param     data  body      request.SkillPackageRequest  true  "工具标识, 技能标识"
// @Success   200   {string}  string                       "打包技能, 返回zip文件流"
// @Router    /skills/packageSkill [post]
func (s *SkillsApi) PackageSkill(c *gin.Context) {
	var req request.SkillPackageRequest
	_ = c.ShouldBindJSON(&req)

	fileName, data, err := skillsService.Package(c.Request.Context(), req)
	if err != nil {
		global.GVA_LOG.Error("打包技能失败", zap.Error(err))
		response.FailWithMessage("打包技能失败: "+err.Error(), c)
		return
	}

	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	c.Data(http.StatusOK, "application/zip", data)
}

// DownloadOnlineSkill
// @Tags      Skills
// @Summary   下载在线技能
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.DownloadOnlineSkillReq  true  "工具标识, 技能ID, 版本"
// @Success   200   {object}  response.Response{msg=string}   "下载在线技能"
// @Router    /skills/downloadOnlineSkill [post]
func (s *SkillsApi) DownloadOnlineSkill(c *gin.Context) {
	var req request.DownloadOnlineSkillReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := skillsService.DownloadOnlineSkill(c.Request.Context(), req); err != nil {
		global.GVA_LOG.Error("下载在线技能失败", zap.Error(err))
		response.FailWithMessage("下载在线技能失败: "+err.Error(), c)
		return
	}
	response.OkWithMessage("下载成功", c)
}
