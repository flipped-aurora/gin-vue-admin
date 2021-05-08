package request

import "gin-vue-admin/model"

type ProductPluginSearch struct{
    model.ProductPlugin
    PageInfo
}