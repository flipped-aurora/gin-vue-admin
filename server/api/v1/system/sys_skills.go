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

func (s *SkillsApi) GetTools(c *gin.Context) {
	data, err := skillsService.Tools(c.Request.Context())
	if err != nil {
		global.GVA_LOG.Error("获取工具列表失败", zap.Error(err))
		response.FailWithMessage("获取工具列表失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"tools": data}, "获取成功", c)
}

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
