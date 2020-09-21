package request

import "gin-vue-admin/model"

type HrpTriagedQueueSearch struct{
    model.HrpTriagedQueue
    PageInfo
}