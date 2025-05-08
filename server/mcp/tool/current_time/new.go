package current_time

// 获取当前系统时间
func (t *Tool) New() (name string, description string, inputReqStruct interface{}) {
	name = "CurrentTime"
	description = "获取当前系统时间"
	inputReqStruct = t
	return
}
