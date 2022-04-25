package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
{{ if .NeedModel }}	"github.com/flipped-aurora/gin-vue-admin/server/plugin/{{ .Snake}}/model" {{ end }}
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/{{ .Snake}}/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type {{ .PlugName}}Api struct{}

// @Tags {{ .PlugName}}
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /{{ .RouterGroup}}/routerName[post]
func (p *{{ .PlugName}}Api) ApiName(c *gin.Context) {
    {{- if .HasRequest}}
        var plug model.Request
        _ = c.ShouldBindJSON(&plug)
    {{ end -}}
        if err{{- if .HasResponse }},res {{ end -}}:= service.ServiceGroupApp.PlugService({{ if .HasRequest }}plug{{ end -}}); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {
	{{if .HasResponse }}
	    response.OkWithDetailed(res,"成功",c)
	{{else}}
	    response.OkWithData("成功", c)
	{{ end -}}

	}
}
