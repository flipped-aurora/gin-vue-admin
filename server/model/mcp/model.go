package model

type McpConfig struct {
	Url     string
	Name    string
	Version string
}

// 获取当前系统时间
type CurrentTimeReq struct {
	Timezone string `json:"timezone" description:"current time timezone"`
}

// 获取天气信息
// 参数：城市名称
// 日期: 2023-07-01
type WeatherReq struct {
	City string `json:"city" description:"city name"`
	Date string `json:"date" description:"date"`
}
