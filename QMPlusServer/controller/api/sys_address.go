package api

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/sysModel"
	"github.com/gin-gonic/gin"
)

func GetProvincesList(c *gin.Context) {
	var pageInfo modelInterface.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	err, list, total := new(sysModel.ChProvinces).GetInfoList(pageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"PrvList":  list,
			"total":    total,
			"page":     pageInfo.Page,
			"pageSize": pageInfo.PageSize,
		})
	}
}

type ProvinceParam struct {
	Province string `json:"province"`
}

func GetCityListByProvince(c *gin.Context) {
	var prv ProvinceParam

	_ = c.ShouldBindJSON(&prv)
	cityList, err := new(sysModel.ChCity).GetCityByProvince(prv.Province)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, false, "获取数据成功", gin.H{"city": cityList})
	}
}
