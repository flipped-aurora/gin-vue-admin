package request

import (
	"github.com/flipped-aurora/gin-vue-admin/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}