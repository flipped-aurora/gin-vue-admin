package request

import "gin-vue-admin/model"

type HrpQueueSearch struct{
    model.HrpQueue
    PageInfo
}