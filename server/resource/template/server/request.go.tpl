package request

import "github.com/eyotang/game-proxy/server/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}