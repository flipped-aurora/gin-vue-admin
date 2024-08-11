{{if .IsPlugin}}
// {{.FuncName}} 等待开发的的{{.Description}}接口
// @Tags {{.StructName}}
// @Summary 等待开发的的{{.Description}}接口
// @accept application/json
// @Produce application/json
// @Param data query request.{{.StructName}}Search true "分页获取{{.Description}}列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /{{.Abbreviation}}/{{.Router}} [{{.Method}}]
export const {{.Abbreviation}}{{.Router}} = () => {
  return service({
    url: '/{{.Abbreviation}}/{{.Router}}',
    method: '{{.Method}}'
  })
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
export const {{.Abbreviation}}{{.Router}} = () => {
  return service({
    url: '/{{.Abbreviation}}/{{.Router}}',
    method: '{{.Method}}'
  })
}

{{- end -}}