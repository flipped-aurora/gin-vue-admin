package tool

import (
	model "github.com/flipped-aurora/gin-vue-admin/server/model/mcp"
)

type ToolList struct{}

// 获取当前系统时间
func (t *ToolList) CurrentTime() (name string, description string, inputReqStruct interface{}) {
	name = "CurrentTime"
	description = "获取当前系统时间"
	inputReqStruct = model.CurrentTimeReq{}
	return
}

// 获取天气
func (t *ToolList) CityWeather() (name string, description string, inputReqStruct interface{}) {
	name = "Weather"
	description = "获取天气"
	inputReqStruct = model.WeatherReq{} //自动解析参数
	return
}
