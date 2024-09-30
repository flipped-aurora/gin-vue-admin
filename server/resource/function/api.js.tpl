{{if .IsPlugin}}
// {{.FuncName}} {{.FuncDesc}}
// @Tags {{.StructName}}
// @Summary {{.FuncDesc}}
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /{{.Abbreviation}}/{{.Router}} [{{.Method}}]
export const {{.Router}} = () => {
  return service({
    url: '/{{.Abbreviation}}/{{.Router}}',
    method: '{{.Method}}'
  })
}

{{- else -}}

// {{.FuncName}} {{.FuncDesc}}
// @Tags {{.StructName}}
// @Summary {{.FuncDesc}}
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /{{.Abbreviation}}/{{.Router}} [{{.Method}}]
export const {{.Router}} = () => {
  return service({
    url: '/{{.Abbreviation}}/{{.Router}}',
    method: '{{.Method}}'
  })
}

{{- end -}}