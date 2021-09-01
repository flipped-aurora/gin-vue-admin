package response

type TextNotify struct { // 文字信息
	Content   string   `json:"content"`   // 发送的内容
	AtMobiles []string `json:"atMobiles"` // 需要艾特的人的手机号
	IsAtAll   bool     `json:"isAtAll"`   // 是否艾特全体
}

type LinkNotify struct { // 图文链接信息
	Content    string `json:"content"`    // 发送的内容
	Title      string `json:"title"`      // 内容标题
	PicUrl     string `json:"picUrl"`     // 配图
	MessageUrl string `json:"messageUrl"` // 点击跳转路径
}

type MarkdownNotify struct { // markdown信息
	Title     string   `json:"title"`     // 内容标题
	Content   string   `json:"content"`   // 发送的内容
	AtMobiles []string `json:"atMobiles"` // 需要艾特的人的手机号
	IsAtAll   bool     `json:"isAtAll"`   // 是否艾特全体
}
