package system

import "github.com/gin-gonic/gin"

type SkillsRouter struct{}

func (s *SkillsRouter) InitSkillsRouter(Router *gin.RouterGroup, pubRouter *gin.RouterGroup) {
	skillsRouter := Router.Group("skills")
	skillsRouterPub := pubRouter.Group("skills")
	{
		skillsRouter.GET("getTools", skillsApi.GetTools)
		skillsRouter.POST("getSkillList", skillsApi.GetSkillList)
		skillsRouter.POST("getSkillDetail", skillsApi.GetSkillDetail)
		skillsRouter.POST("saveSkill", skillsApi.SaveSkill)
		skillsRouter.POST("deleteSkill", skillsApi.DeleteSkill)
		skillsRouter.POST("createScript", skillsApi.CreateScript)
		skillsRouter.POST("getScript", skillsApi.GetScript)
		skillsRouter.POST("saveScript", skillsApi.SaveScript)
		skillsRouter.POST("createResource", skillsApi.CreateResource)
		skillsRouter.POST("getResource", skillsApi.GetResource)
		skillsRouter.POST("saveResource", skillsApi.SaveResource)
		skillsRouter.POST("createReference", skillsApi.CreateReference)
		skillsRouter.POST("getReference", skillsApi.GetReference)
		skillsRouter.POST("saveReference", skillsApi.SaveReference)
		skillsRouter.POST("createTemplate", skillsApi.CreateTemplate)
		skillsRouter.POST("getTemplate", skillsApi.GetTemplate)
		skillsRouter.POST("saveTemplate", skillsApi.SaveTemplate)
		skillsRouter.POST("getGlobalConstraint", skillsApi.GetGlobalConstraint)
		skillsRouter.POST("saveGlobalConstraint", skillsApi.SaveGlobalConstraint)
		skillsRouter.POST("packageSkill", skillsApi.PackageSkill)
	}
	{
		skillsRouterPub.POST("downloadOnlineSkill", skillsApi.DownloadOnlineSkill)
	}
}
