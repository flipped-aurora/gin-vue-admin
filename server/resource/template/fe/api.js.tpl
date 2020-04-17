import service from '@/utils/request'

// @Tags {{.StructName}}
// @Summary 创建{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.PackageName}}.{{.StructName}} true "创建{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /{{.Abbreviation}}/create{{.StructName}} [post]
export const create{{.StructName}} = (data) => {
     return service({
         url: "/{{.Abbreviation}}/create{{.StructName}}",
         method: 'post',
         data
     })
 }


// @Tags {{.StructName}}
// @Summary 删除{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.PackageName}}.{{.StructName}} true "删除{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /{{.Abbreviation}}/delete{{.StructName}} [post]
 export const delete{{.StructName}} = (data) => {
     return service({
         url: "/{{.Abbreviation}}/delete{{.StructName}}",
         method: 'post',
         data
     })
 }

// @Tags {{.StructName}}
// @Summary 更新{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.PackageName}}.{{.StructName}} true "更新{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /{{.Abbreviation}}/update{{.StructName}} [post]
 export const update{{.StructName}} = (data) => {
     return service({
         url: "/{{.Abbreviation}}/update{{.StructName}}",
         method: 'post',
         data
     })
 }


// @Tags {{.StructName}}
// @Summary 用id查询{{.StructName}}
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {{.PackageName}}.{{.StructName}} true "用id查询{{.StructName}}"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /{{.Abbreviation}}/find{{.StructName}} [post]
 export const find{{.StructName}} = (data) => {
     return service({
         url: "/{{.Abbreviation}}/find{{.StructName}}",
         method: 'post',
         data
     })
 }


// @Tags {{.StructName}}
// @Summary 分页获取{{.StructName}}列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取{{.StructName}}列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /{{.Abbreviation}}/get{{.StructName}}List [post]
 export const get{{.StructName}}List = (data) => {
     return service({
         url: "/{{.Abbreviation}}/get{{.StructName}}List",
         method: 'post',
         data
     })
 }