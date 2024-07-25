// Get{{.StructName}}Public 等待开发的的{{.Description}}接口
// @Tags {{.StructName}}
// @Summary 不需要鉴权的{{.Description}}接口
// @accept application/json
// @Produce application/json
// @Param data query {{.Package}}Req.{{.StructName}}Search true "分页获取{{.Description}}列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /{{.Abbreviation}}/{{.FuncName}} [{{.Method}}]
func ({{.Abbreviation}}Api *{{.StructName}}Api) {{.FuncName}}(c *gin.Context) {
// 请添加自己的业务逻辑
if dataSource, err := {{.Abbreviation}}Service.{{.FuncName}}(); err != nil {
        global.GVA_LOG.Error("失败!", zap.Error(err))
   		response.FailWithMessage("失败", c)
   	} else {
   		response.OkWithData("返回数据",c)
   	}
}