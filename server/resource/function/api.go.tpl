{{if .IsPlugin}}
// {{.FuncName}} 等待开发的的{{.Description}}接口
// @Tags {{.StructName}}
// @Summary 等待开发的的{{.Description}}接口
// @accept application/json
// @Produce application/json
// @Param data query request.{{.StructName}}Search true "分页获取{{.Description}}列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /{{.Abbreviation}}/{{.FuncName}} [{{.Method}}]
func (a *{{.Abbreviation}}) {{.Router}}(c *gin.Context) {
    // 请添加自己的业务逻辑
    err := service{{ .StructName }}.{{.FuncName}}()
       if err != nil {
    		global.GVA_LOG.Error("失败!", zap.Error(err))
            response.FailWithMessage("失败", c)
    		return
       }
    response.OkWithData("返回数据",c)
}

{{- else -}}

// {{.FuncName}} 等待开发的的{{.Description}}接口
// @Tags {{.StructName}}
// @Summary 等待开发的的{{.Description}}接口
// @accept application/json
// @Produce application/json
// @Param data query {{.Package}}Req.{{.StructName}}Search true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /{{.Abbreviation}}/{{.Router}} [{{.Method}}]
func ({{.Abbreviation}}Api *{{.StructName}}Api){{.FuncName}}(c *gin.Context) {
    // 请添加自己的业务逻辑
    if err := {{.Abbreviation}}Service.{{.FuncName}}()
    if err != nil {
        global.GVA_LOG.Error("失败!", zap.Error(err))
   		response.FailWithMessage("失败", c)
   		return
   	}
   	response.OkWithData("返回数据",c)
}
{{end}}