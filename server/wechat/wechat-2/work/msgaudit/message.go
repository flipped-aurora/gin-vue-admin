package msgaudit

// BaseMessage 基础消息
type BaseMessage struct {
	MsgID   string   `json:"msgid,omitempty"`   // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	Action  string   `json:"action,omitempty"`  // 消息动作，目前有send(发送消息)/recall(撤回消息)/switch(切换企业日志)三种类型。
	From    string   `json:"from,omitempty"`    // 消息发送方id。同一企业内容为userid，非相同企业为external_userid。消息如果是机器人发出，也为external_userid。
	ToList  []string `json:"tolist,omitempty"`  // 消息接收方列表，可能是多个，同一个企业内容为userid，非相同企业为external_userid。
	RoomID  string   `json:"roomid,omitempty"`  // 群聊消息的群id。如果是单聊则为空。
	MsgTime int64    `json:"msgtime,omitempty"` // 消息发送时间戳，utc时间，ms单位。
	MsgType string   `json:"msgtype,omitempty"` // 文本消息为：text。
}

// TextMessage 文本消息
type TextMessage struct {
	BaseMessage
	Text struct {
		Content string `json:"content,omitempty"` // 消息内容。
	} `json:"text,omitempty"`
}

// ImageMessage 图片消息
type ImageMessage struct {
	BaseMessage
	Image struct {
		SdkFileID string `json:"sdkfileid,omitempty"` // 媒体资源的id信息。
		Md5Sum    string `json:"md5sum,omitempty"`    // 图片资源的md5值，供进行校验。
		FileSize  uint32 `json:"filesize,omitempty"`  // 图片资源的文件大小。
	} `json:"image,omitempty"`
}

// RevokeMessage 撤回消息
type RevokeMessage struct {
	BaseMessage
	Revoke struct {
		PreMsgID string `json:"pre_msgid,omitempty"` // 标识撤回的原消息的msgid
	} `json:"revoke,omitempty"`
}

// AgreeMessage 同意会话聊天内容
type AgreeMessage struct {
	BaseMessage
	Agree struct {
		UserID    string `json:"userid,omitempty"`     // 同意/不同意协议者的userid，外部企业默认为external_userid。
		AgreeTime int64  `json:"agree_time,omitempty"` // 同意/不同意协议的时间，utc时间，ms单位。
	} `json:"agree,omitempty"`
}

// VoiceMessage 语音消息
type VoiceMessage struct {
	BaseMessage
	Voice struct {
		SdkFileID  string `json:"sdkfileid,omitempty"`   // 媒体资源的id信息。
		VoiceSize  uint32 `json:"voice_size,omitempty"`  // 语音消息大小。
		PlayLength uint32 `json:"play_length,omitempty"` // 播放长度。
		Md5Sum     string `json:"md5sum,omitempty"`      // 图片资源的md5值，供进行校验。
	} `json:"voice,omitempty"`
}

// VideoMessage 视频消息
type VideoMessage struct {
	BaseMessage
	Video struct {
		SdkFileID  string `json:"sdkfileid,omitempty"`   // 媒体资源的id信息。
		FileSize   uint32 `json:"filesize,omitempty"`    // 图片资源的文件大小。
		PlayLength uint32 `json:"play_length,omitempty"` // 播放长度。
		Md5Sum     string `json:"md5sum,omitempty"`      // 图片资源的md5值，供进行校验。
	} `json:"video,omitempty"`
}

// CardMessage 名片消息
type CardMessage struct {
	BaseMessage
	Card struct {
		CorpName string `json:"corpname,omitempty"` // 名片所有者所在的公司名称。
		UserID   string `json:"userid,omitempty"`   // 名片所有者的id，同一公司是userid，不同公司是external_userid
	} `json:"card,omitempty"`
}

// LocationMessage 位置消息
type LocationMessage struct {
	BaseMessage
	Location struct {
		Lng     float64 `json:"longitude,omitempty"` // 经度，单位double
		Lat     float64 `json:"latitude,omitempty"`  // 纬度，单位double
		Address string  `json:"address,omitempty"`   // 地址信息
		Title   string  `json:"title,omitempty"`     // 位置信息的title。
		Zoom    uint32  `json:"zoom,omitempty"`      // 缩放比例。
	} `json:"location,omitempty"`
}

// EmotionMessage 表情消息
type EmotionMessage struct {
	BaseMessage
	Emotion struct {
		Type      uint32 `json:"type,omitempty"`      // 表情类型，png或者gif.1表示gif 2表示png。
		Width     uint32 `json:"width,omitempty"`     // 表情图片宽度。
		Height    uint32 `json:"height,omitempty"`    // 表情图片高度。
		ImageSize uint32 `json:"imagesize,omitempty"` // 资源的文件大小。
		SdkFileID string `json:"sdkfileid,omitempty"` // 媒体资源的id信息。
		Md5Sum    string `json:"md5sum,omitempty"`    // 图片资源的md5值，供进行校验。
	} `json:"emotion,omitempty"`
}

// FileMessage 文件消息
type FileMessage struct {
	BaseMessage
	File struct {
		FileName  string `json:"filename,omitempty"`  // 文件名称。
		FileExt   string `json:"fileext,omitempty"`   // 文件类型后缀。
		SdkFileID string `json:"sdkfileid,omitempty"` // 媒体资源的id信息。
		FileSize  uint32 `json:"filesize,omitempty"`  // 文件大小。
		Md5Sum    string `json:"md5sum,omitempty"`    // 资源的md5值，供进行校验。
	} `json:"file,omitempty"`
}

// LinkMessage 链接消息
type LinkMessage struct {
	BaseMessage
	Link struct {
		Title    string `json:"title,omitempty"`       // 消息标题。
		Desc     string `json:"description,omitempty"` // 消息描述。
		LinkURL  string `json:"link_url,omitempty"`    // 链接url地址
		ImageURL string `json:"image_url,omitempty"`   // 链接图片url。
	} `json:"link,omitempty"`
}

// WeappMessage 小程序消息
type WeappMessage struct {
	BaseMessage
	WeApp struct {
		Title       string `json:"title,omitempty"`       // 消息标题。
		Desc        string `json:"description,omitempty"` // 消息描述。
		Username    string `json:"username,omitempty"`    // 用户名称。
		DisplayName string `json:"displayname,omitempty"` // 小程序名称
	} `json:"weapp,omitempty"`
}

// ChatRecordMessage 会话记录消息
type ChatRecordMessage struct {
	BaseMessage
	ChatRecord struct {
		Title string       `json:"title,omitempty"` // 聊天记录标题
		Item  []ChatRecord `json:"item,omitempty"`  // 消息记录内的消息内容，批量数据
	} `json:"chatrecord,omitempty"`
}

// TodoMessage 待办消息
type TodoMessage struct {
	BaseMessage
	Todo struct {
		Title   string `json:"title,omitempty"`   // 代办的来源文本
		Content string `json:"content,omitempty"` // 	代办的具体内容
	} `json:"todo,omitempty"`
}

// VoteMessage 投票消息
type VoteMessage struct {
	BaseMessage
	VoteTitle string   `json:"votetitle,omitempty"` // 投票主题。
	VoteItem  []string `json:"voteitem,omitempty"`  // 投票选项，可能多个内容。
	VoteType  uint32   `json:"votetype,omitempty"`  // 投票类型.101发起投票、102参与投票。
	VoteID    string   `json:"voteid,omitempty"`    // 投票id，方便将参与投票消息与发起投票消息进行前后对照。
}

// CollectMessage 填表消息
type CollectMessage struct {
	BaseMessage
	Collect struct {
		RoomName   string           `json:"room_name,omitempty"`   // 填表消息所在的群名称。
		Creator    string           `json:"creator,omitempty"`     // 创建者在群中的名字
		CreateTime string           `json:"create_time,omitempty"` // 创建的时间
		Details    []CollectDetails `json:"details,omitempty"`     // 表内容
	} `json:"collect,omitempty"`
}

// RedpacketMessage 红包消息
type RedpacketMessage struct {
	BaseMessage
	RedPacket struct {
		Type        uint32 `json:"type,omitempty"`        // 红包消息类型。1 普通红包、2 拼手气群红包、3 激励群红包。
		Wish        string `json:"wish,omitempty"`        // 红包祝福语
		TotalCnt    uint32 `json:"totalcnt,omitempty"`    // 红包总个数
		TotalAmount uint32 `json:"totalamount,omitempty"` // 红包总金额。单位为分。
	} `json:"redpacket,omitempty"`
}

// MeetingMessage 会议邀请消息
type MeetingMessage struct {
	BaseMessage
	Meeting struct {
		Topic       string `json:"topic,omitempty"`       // 会议主题
		StartTime   int64  `json:"starttime,omitempty"`   // 会议开始时间。Utc时间
		EndTime     int64  `json:"endtime,omitempty"`     // 会议结束时间。Utc时间
		Address     string `json:"address,omitempty"`     // 会议地址
		Remarks     string `json:"remarks,omitempty"`     // 会议备注
		MeetingType uint32 `json:"meetingtype,omitempty"` // 会议消息类型。101发起会议邀请消息、102处理会议邀请消息
		MeetingID   uint64 `json:"meetingid,omitempty"`   // 会议id。方便将发起、处理消息进行对照
		Status      uint32 `json:"status,omitempty"`      // 会议邀请处理状态。1 参加会议、2 拒绝会议、3 待定、4 未被邀请、5 会议已取消、6 会议已过期、7 不在房间内。
	} `json:"meeting,omitempty"`
}

// DocMessage 在线文档消息
type DocMessage struct {
	BaseMessage
	Doc struct {
		Title      string `json:"title,omitempty"`       // 在线文档名称
		LinkURL    string `json:"link_url,omitempty"`    // 在线文档链接
		DocCreator string `json:"doc_creator,omitempty"` // 在线文档创建者。本企业成员创建为userid；外部企业成员创建为external_userid
	} `json:"doc,omitempty"`
}

// MarkdownMessage MarkDown消息
type MarkdownMessage struct {
	BaseMessage
	Info struct {
		Content string `json:"content,omitempty"` // markdown消息内容，目前为机器人发出的消息
	} `json:"info,omitempty"`
}

// NewsMessage 图文消息
type NewsMessage struct {
	BaseMessage
	Info struct {
		Item []News `json:"item,omitempty"` // 图文消息数组
	} `json:"info,omitempty"` // 图文消息的内容
}

// CalendarMessage 日程消息
type CalendarMessage struct {
	BaseMessage
	Calendar struct {
		Title        string   `json:"title,omitempty"`        // 日程主题
		CreatorName  string   `json:"creatorname,omitempty"`  // 日程组织者
		AttendeeName []string `json:"attendeename,omitempty"` // 日程参与人。数组，内容为String类型
		StartTime    int64    `json:"starttime,omitempty"`    // 日程开始时间。Utc时间，单位秒
		EndTime      int64    `json:"endtime,omitempty"`      // 日程结束时间。Utc时间，单位秒
		Place        string   `json:"place,omitempty"`        // 日程地点
		Remarks      string   `json:"remarks,omitempty"`      // 日程备注
	} `json:"calendar,omitempty"`
}

// MixedMessage 混合消息
type MixedMessage struct {
	BaseMessage
	Mixed struct {
		Item []MixedMsg `json:"item,omitempty"`
	} `json:"mixed,omitempty"` // 消息内容。可包含图片、文字、表情等多种消息。Object类型
}

// MeetingVoiceCallMessage 音频存档消息
type MeetingVoiceCallMessage struct {
	BaseMessage
	VoiceID          string            `json:"voiceid,omitempty"`            // 音频id
	MeetingVoiceCall *MeetingVoiceCall `json:"meeting_voice_call,omitempty"` // 音频消息内容。包括结束时间、fileid，可能包括多个demofiledata、sharescreendata消息，demofiledata表示文档共享信息，sharescreendata表示屏幕共享信息。Object类型
}

// VoipDocShareMessage 音频共享消息
type VoipDocShareMessage struct {
	BaseMessage
	VoipID       string        `json:"voipid,omitempty"`         // 音频id
	VoipDocShare *VoipDocShare `json:"voip_doc_share,omitempty"` // 共享文档消息内容。包括filename、md5sum、filesize、sdkfileid字段。Object类型
}

// ExternalRedPacketMessage 互通小红包消息
type ExternalRedPacketMessage struct {
	BaseMessage
	RedPacket struct {
		Type        int32 `json:"type,omitempty"`        // 红包消息类型。1 普通红包、2 拼手气群红包。Uint32类型
		Wish        int32 `json:"wish,omitempty"`        // 红包祝福语。String类型
		TotalCnt    int32 `json:"totalcnt,omitempty"`    // 红包总个数。Uint32类型
		TotalAmount int32 `json:"totalamount,omitempty"` // 红包消息类型。1 普通红包、2 拼手气群红包。Uint32类型
	} `json:"redpacket,omitempty"`
}

// SphFeedMessage 视频号消息
type SphFeedMessage struct {
	BaseMessage
	SphFeed struct {
		FeedType string `json:"feed_type,omitempty"` // 视频号消息类型
		SphName  string `json:"sph_name,omitempty"`  // 视频号账号名称
		FeedDesc uint64 `json:"feed_desc,omitempty"` // 视频号账号名称
	}
}

// SwitchMessage 企业切换日志
type SwitchMessage struct {
	MsgID  string `json:"msgid,omitempty"`  // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重
	Action string `json:"action,omitempty"` // 消息动作，切换企业为switch
	Time   int64  `json:"time,omitempty"`   // 消息发送时间戳，utc时间，ms单位。
	User   string `json:"user,omitempty"`   // 具体为切换企业的成员的userid。
}

// ChatRecord 会话记录消息
type ChatRecord struct {
	Type         string `json:"type,omitempty"`          // 每条聊天记录的具体消息类型：ChatRecordText/ ChatRecordFile/ ChatRecordImage/ ChatRecordVideo/ ChatRecordLink/ ChatRecordLocation/ ChatRecordMixed ….
	Content      string `json:"content,omitempty"`       // 消息内容。Json串，内容为对应类型的json
	MsgTime      int64  `json:"msgtime,omitempty"`       // 消息时间，utc时间，ms单位。
	FromChatroom bool   `json:"from_chatroom,omitempty"` // 是否来自群会话。
}

// CollectDetails 填表消息
type CollectDetails struct {
	ID   uint64 `json:"id,omitempty"`   // 表项id
	Ques string `json:"ques,omitempty"` // 表项名称
	Type string `json:"type,omitempty"` // 表项类型，有Text(文本),Number(数字),Date(日期),Time(时间)
}

// News 图文消息
type News struct {
	Title  string `json:"title,omitempty"`       // 图文消息标题
	Desc   string `json:"description,omitempty"` // 图文消息描述
	URL    string `json:"url,omitempty"`         // 图文消息点击跳转地址
	PicURL string `json:"picurl,omitempty"`      // 图文消息配图的url
}

// MixedMsg 混合消息
type MixedMsg struct {
	Type    string `json:"type,omitempty"`
	Content string `json:"content,omitempty"`
}

// MeetingVoiceCall 音频存档消息
type MeetingVoiceCall struct {
	EndTime         int64             `json:"endtime,omitempty"`         // 音频结束时间
	SdkFileID       string            `json:"sdkfileid,omitempty"`       // 音频媒体下载的id
	DemoFileData    []DemoFileData    `json:"demofiledata,omitempty"`    // 文档分享对象，Object类型
	ShareScreenData []ShareScreenData `json:"sharescreendata,omitempty"` // 屏幕共享对象，Object类型
}

// DemoFileData 文档共享消息
type DemoFileData struct {
	FileName     string `json:"filename,omitempty"`     // 文档共享名称
	DemoOperator string `json:"demooperator,omitempty"` // 文档共享操作用户的id
	StartTime    int64  `json:"starttime,omitempty"`    // 文档共享开始时间
	EndTime      int64  `json:"endtime,omitempty"`      // 文档共享结束时间
}

// ShareScreenData 屏幕共享信息
type ShareScreenData struct {
	Share     string `json:"share,omitempty"`     // 屏幕共享用户的id
	StartTime int64  `json:"starttime,omitempty"` // 屏幕共享开始时间
	EndTime   int64  `json:"endtime,omitempty"`   // 屏幕共享结束时间
}

// VoipDocShare 音频共享文档消息
type VoipDocShare struct {
	FileName  string `json:"filename,omitempty"`  // 文档共享文件名称
	Md5Sum    string `json:"md5sum,omitempty"`    // 共享文件的md5值
	FileSize  uint64 `json:"filesize,omitempty"`  // 共享文件的大小
	SdkFileID string `json:"sdkfileid,omitempty"` // 共享文件的sdkfile，通过此字段进行媒体数据下载
}
