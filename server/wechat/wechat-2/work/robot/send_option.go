package robot

import "github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/util"

// WebhookSendResponse 机器人发送群组消息响应
type WebhookSendResponse struct {
	util.CommonError
}

// WebhookSendTextOption 机器人发送文本消息请求参数
type WebhookSendTextOption struct {
	MsgType string `json:"msgtype"` // 消息类型,此时固定为text
	Text    struct {
		Content             string   `json:"content"`               // 文本内容，最长不超过2048个字节，必须是utf8编码
		MentionedList       []string `json:"mentioned_list"`        // userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
		MentionedMobileList []string `json:"mentioned_mobile_list"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
	} `json:"text"` // 文本消息内容
}

// WebhookSendMarkdownOption 机器人发送markdown消息请求参数
// 支持语法参考 https://developer.work.weixin.qq.com/document/path/91770
type WebhookSendMarkdownOption struct {
	MsgType  string `json:"msgtype"` // 消息类型,此时固定为markdown
	Markdown struct {
		Content string `json:"content"` // markdown内容，最长不超过4096个字节，必须是utf8编码
	} `json:"markdown"` // markdown消息内容
}

// WebhookSendImageOption 机器人发送图片消息请求参数
type WebhookSendImageOption struct {
	MsgType string `json:"msgtype"` // 消息类型,此时固定为image
	Image   struct {
		Base64 string `json:"base64"` // 图片内容的base64编码
		MD5    string `json:"md5"`    // 图片内容（base64编码前）的md5值
	} `json:"image"` // 图片消息内容
}

// WebhookSendNewsOption 机器人发送图文消息请求参数
type WebhookSendNewsOption struct {
	MsgType string `json:"msgtype"` // 消息类型,此时固定为news
	News    struct {
		Articles []struct {
			Title       string `json:"title"`       // 标题，不超过128个字节，超过会自动截断
			Description string `json:"description"` // 描述，不超过512个字节，超过会自动截断
			URL         string `json:"url"`         // 点击后跳转的链接
			PicURL      string `json:"picurl"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150
		} `json:"articles"` // 图文消息列表 一个图文消息支持1到8条图文
	} `json:"news"` // 图文消息内容
}

// WebhookSendFileOption 机器人发送文件消息请求参数
type WebhookSendFileOption struct {
	MsgType string `json:"msgtype"` // 消息类型，此时固定为file
	File    struct {
		MediaID string `json:"media_id"` // 文件id，通过下文的文件上传接口获取
	} `json:"file"` // 文件类型
}

// WebHookSendTempNoticeOption 机器人发送文本通知模版消息请求参数
type WebHookSendTempNoticeOption struct {
	MsgType      string       `json:"msgtype"`       // 消息类型，此时的消息类型固定为template_card
	TemplateCard TemplateCard `json:"template_card"` // 具体的模版卡片参数
}

// TemplateCard 具体的模版卡片参数
type TemplateCard struct {
	CardType              string        `json:"card_type"`               // 模版卡片的模版类型，文本通知模版卡片的类型为text_notice
	Source                CardSource    `json:"source"`                  // 卡片来源样式信息，不需要来源样式可不填写
	MainTitle             CardTitle     `json:"main_title"`              // 模版卡片的主要内容，包括一级标题和标题辅助信息
	EmphasisContent       CardTitle     `json:"emphasis_content"`        // 关键数据样式
	QuoteArea             CardQuoteArea `json:"quote_area"`              // 引用文献样式，建议不与关键数据共用
	SubTitleText          string        `json:"sub_title_text"`          // 二级普通文本，建议不超过112个字。模版卡片主要内容的一级标题main_title.title和二级普通文本sub_title_text必须有一项填写
	HorizontalContentList []CardContent `json:"horizontal_content_list"` // 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
	JumpList              []JumpContent `json:"jump_list"`               // 跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
	CardAction            CardAction    `json:"card_action"`             // 整体卡片的点击跳转事件，text_notice模版卡片中该字段为必填项
}

// CardSource 卡片来源样式信息，不需要来源样式可不填写
type CardSource struct {
	IconURL   string `json:"icon_url"`   // 来源图片的url
	Desc      string `json:"desc"`       // 来源图片的描述，建议不超过13个字
	DescColor int    `json:"desc_color"` // 来源文字的颜色，目前支持：0(默认) 灰色，1 黑色，2 红色，3 绿色
}

// CardTitle 标题和标题辅助信息
type CardTitle struct {
	Title string `json:"title"` // 标题，建议不超过26个字。模版卡片主要内容的一级标题main_title.title和二级普通文本sub_title_text必须有一项填写
	Desc  string `json:"desc"`  // 标题辅助信息，建议不超过30个字
}

// CardQuoteArea 引用文献样式，建议不与关键数据共用
type CardQuoteArea struct {
	Type      int    `json:"type"`               // 引用文献样式区域点击事件，0或不填代表没有点击事件，1 代表跳转url，2 代表跳转小程序
	URL       string `json:"url,omitempty"`      // 点击跳转的url，quote_area.type是1时必填
	Appid     string `json:"appid,omitempty"`    // 点击跳转的小程序的appid，quote_area.type是2时必填
	Pagepath  string `json:"pagepath,omitempty"` // 点击跳转的小程序的pagepath，quote_area.type是2时选填
	Title     string `json:"title"`              // 引用文献样式的标题
	QuoteText string `json:"quote_text"`         // 引用文献样式的引用文案
}

// CardContent 二级标题+文本列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过6
type CardContent struct {
	KeyName string `json:"keyname"`            // 链接类型，0或不填代表是普通文本，1 代表跳转url，2 代表下载附件，3 代表@员工
	Value   string `json:"value"`              // 二级标题，建议不超过5个字
	Type    int    `json:"type,omitempty"`     // 二级文本，如果horizontal_content_list.type是2，该字段代表文件名称（要包含文件类型），建议不超过26个字
	URL     string `json:"url,omitempty"`      // 链接跳转的url，horizontal_content_list.type是1时必填
	MediaID string `json:"media_id,omitempty"` // 附件的media_id，horizontal_content_list.type是2时必填
	UserID  string `json:"userid,omitempty"`   // 被@的成员的userid，horizontal_content_list.type是3时必填
}

// JumpContent 跳转指引样式的列表，该字段可为空数组，但有数据的话需确认对应字段是否必填，列表长度不超过3
type JumpContent struct {
	Type     int    `json:"type"`               // 跳转链接类型，0或不填代表不是链接，1 代表跳转url，2 代表跳转小程序
	URL      string `json:"url,omitempty"`      // 跳转链接的url，jump_list.type是1时必填
	Title    string `json:"title"`              // 跳转链接样式的文案内容，建议不超过13个字
	AppID    string `json:"appid,omitempty"`    // 跳转链接的小程序的appid，jump_list.type是2时必填
	PagePath string `json:"pagepath,omitempty"` // 跳转链接的小程序的pagepath，jump_list.type是2时选填
}

// CardAction 整体卡片的点击跳转事件，text_notice模版卡片中该字段为必填项
type CardAction struct {
	Type     int    `json:"type"`               // 卡片跳转类型，1 代表跳转url，2 代表打开小程序。text_notice模版卡片中该字段取值范围为[1,2]
	URL      string `json:"url,omitempty"`      // 跳转事件的url，card_action.type是1时必填
	Appid    string `json:"appid,omitempty"`    // 跳转事件的小程序的appid，card_action.type是2时必填
	PagePath string `json:"pagepath,omitempty"` // 跳转事件的小程序的pagepath，card_action.type是2时选填
}
