package request

import "gin-vue-admin/model"

type SysDictionaryDetailSearch struct{
    model.SysDictionaryDetail
    PageInfo
}