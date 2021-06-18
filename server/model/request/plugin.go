package request

import "github.com/eyotang/game-proxy/server/model"

type ProductPluginSearch struct{
    model.ProductPlugin
    PageInfo
}