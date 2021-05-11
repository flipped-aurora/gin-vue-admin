package request

import "github.com/eyotang/game-api-admin/server/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}