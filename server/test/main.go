package main

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
	"github.com/liu-cn/json-filter/filter"
)

func main() {
	m := filter.SelectMarshal("rollbackVersion", &biz_apphub.BizAppHubRecord{}).Map()
	fmt.Println(m)
}
