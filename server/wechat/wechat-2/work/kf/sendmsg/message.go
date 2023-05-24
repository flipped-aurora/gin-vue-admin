package sendmsg

// Message 发送消息
type Message struct {
	ToUser   string `json:"touser"`          // 指定接收消息的客户UserID
	OpenKFID string `json:"open_kfid"`       // 指定发送消息的客服帐号ID
	MsgID    string `json:"msgid,omitempty"` // 指定消息ID
}

// Text 发送文本消息
type Text struct {
	Message
	MsgType string `json:"msgtype"` // 消息类型，此时固定为：text
	Text    struct {
		Content string `json:"content"` // 消息内容，最长不超过2048个字节
	} `json:"text"` // 文本消息
}

// Image 发送图片消息
type Image struct {
	Message
	MsgType string `json:"msgtype"` // 消息类型，此时固定为：image
	Image   struct {
		MediaID string `json:"media_id"` // 图片文件id，可以调用上传临时素材接口获取
	} `json:"image"` // 图片消息
}

// Voice 发送语音消息
type Voice struct {
	Message
	MsgType string `json:"msgtype"` // 消息类型，此时固定为：voice
	Voice   struct {
		MediaID string `json:"media_id"` // 语音文件id，可以调用上传临时素材接口获取
	} `json:"voice"` // 语音消息
}

// Video 发送视频消息
type Video struct {
	Message
	MsgType string `json:"msgtype"` // 消息类型，此时固定为：video
	Video   struct {
		MediaID string `json:"media_id"` // 视频文件id，可以调用上传临时素材接口获取
	} `json:"video"` // 视频消息
}

// File 发送文件消息
type File struct {
	Message
	MsgType string `json:"msgtype"` // 消息类型，此时固定为：file
	File    struct {
		MediaID string `json:"media_id"` // 文件id，可以调用上传临时素材接口获取
	} `json:"file"` // 文件消息
}

// Link 图文链接消息
type Link struct {
	Message
	MsgType string `json:"msgtype"` // 消息类型，此时固定为：link
	Link    struct {
		Title        string `json:"title"`          // 标题，不超过128个字节，超过会自动截断
		Desc         string `json:"desc"`           // 描述，不超过512个字节，超过会自动截断
		URL          string `json:"url"`            // 点击后跳转的链接。 最长2048字节，请确保包含了协议头(http/https)
		ThumbMediaID string `json:"thumb_media_id"` // 缩略图的media_id, 可以通过素材管理接口获得。此处thumb_media_id即上传接口返回的media_id
	} `json:"link"` // 链接消息
}

// MiniProgram 小程序消息
type MiniProgram struct {
	Message
	MsgType     string `json:"msgtype"` // 消息类型，此时固定为：miniprogram
	MiniProgram struct {
		AppID        string `json:"appid"`          // 小程序appid，必须是关联到企业的小程序应用
		Title        string `json:"title"`          // 小程序消息标题，最多64个字节，超过会自动截断
		ThumbMediaID string `json:"thumb_media_id"` // 小程序消息封面的mediaid，封面图建议尺寸为520*416
		PagePath     string `json:"pagepath"`       // 点击消息卡片后进入的小程序页面路径
	} `json:"miniprogram"` // 小程序消息
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

// Location 地理位置消息
type Location struct {
	Message
	MsgType  string `json:"msgtype"` // 消息类型，此时固定为：location
	Location struct {
		Latitude  float32 `json:"latitude"`  // 纬度, 浮点数，范围为90 ~ -90
		Longitude float32 `json:"longitude"` // 经度, 浮点数，范围为180 ~ -180
		Name      string  `json:"name"`      // 位置名
		Address   string  `json:"address"`   // 地址详情说明
	} `json:"location"`
}
