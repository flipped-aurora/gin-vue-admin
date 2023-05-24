package msgaudit

// Config 会话存档初始化参数
type Config struct {
	CorpID        string // 调用企业的企业id，例如：wwd08c8exxxx5ab44d，可以在企业微信管理端--我的企业--企业信息查看
	CorpSecret    string // 聊天内容存档的Secret，可以在企业微信管理端--管理工具--聊天内容存档查看
	RasPrivateKey string // 消息加密私钥，可以在企业微信管理端--管理工具--消息加密公钥查看对用公钥，私钥一般由自己保存
}
