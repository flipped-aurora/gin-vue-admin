package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/{{.Package}}"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
