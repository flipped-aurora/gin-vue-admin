package request

import "gin-vue-admin/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}