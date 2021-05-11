package request

import "github.com/eyotang/game-api-admin/server/model"

type ProductPluginSearch struct{
    model.ProductPlugin
    PageInfo
}