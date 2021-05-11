package request

import "github.com/eyotang/gin-vue-admin/server/model"

type ProductPluginSearch struct{
    model.ProductPlugin
    PageInfo
}