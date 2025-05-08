package weather

// 获取天气
func (t *Tool) New() (name string, description string, inputReqStruct interface{}) {
	name = "Weather"
	description = "获取天气"
	inputReqStruct = t //自动解析参数
	return
}
