package sendmsgonevent

// Message 发送事件响应消息
type Message struct {
	Code  string `json:"code"`  // 事件响应消息对应的code。通过事件回调下发，仅可使用一次。
	MsgID string `json:"msgid"` // 消息ID。如果请求参数指定了msgid，则原样返回，否则系统自动生成并返回。不多于32字节，不多于32字节
}

// Text 文本消息
type Text struct {
	Message
	MsgType string `json:"msgtype"` // 消息类型，此时固定为：text
	Text    struct {
		Content string `json:"content"` // 消息内容，最长不超过2048个字节
	} `json:"text"` // 文本消息
}

// Menu 发送菜单消息
type Menu struct {
	Message
	MsgType string `json:"msgtype"` // 消息类型，此时固定为：msgmenu
	MsgMenu struct {
		HeadContent string        `json:"head_content"` // 消息内容，不多于1024字节
		List        []interface{} `json:"list"`         // 菜单项配置，不能多余10个
		TailContent string        `json:"tail_content"` // 结束文本, 不多于1024字
	} `json:"msgmenu"`
}

// MenuClick 回复菜单
type MenuClick struct {
	Type  string `json:"type"` // 菜单类型: click 回复菜单
	Click struct {
		ID      string `json:"id"`      // 菜单ID, 不少于1字节, 不多于64字节
		Content string `json:"content"` // 菜单显示内容, 不少于1字节, 不多于128字节
	} `json:"click"`
}

// MenuView 超链接菜单
type MenuView struct {
	Type string `json:"type"` // 菜单类型: view 超链接菜单
	View struct {
		URL     string `json:"url"`     // 点击后跳转的链接, 不少于1字节, 不多于2048字节
		Content string `json:"content"` // 菜单显示内容, 不少于1字节, 不多于1024字节
	} `json:"view"`
}

// MenuMiniProgram 小程序菜单
type MenuMiniProgram struct {
	Type        string `json:"type"` // 菜单类型: miniprogram 小程序菜单
	MiniProgram struct {
		AppID    string `json:"appid"`    // 小程序appid, 不少于1字节, 不多于32字节
		PagePath string `json:"pagepath"` // 点击后进入的小程序页面, 不少于1字节, 不多于1024字节
		Content  string `json:"content"`  // 菜单显示内容, 不少于1字节, 不多于1024字节
	} `json:"miniprogram"`
}
