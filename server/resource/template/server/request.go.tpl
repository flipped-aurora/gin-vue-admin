package request

import (
	"gin-vue-admin/model/autocode"
	"gin-vue-admin/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}