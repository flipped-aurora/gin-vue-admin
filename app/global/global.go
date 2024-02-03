package global

// mysql配置
const DSN = ""

// JWT配置
const JWTKey = "ab_project"

// 小程序配置
const APPID = ""
const SECRET = ""

// 用户身份
const UserModelPuTong = 0     //普通用户
const UserModelGuanLiYuan = 9 //超级用户
//帖子模型

type aliyunOSSConfig struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	BucketUrl       string
	BasePath        string
}

var AliyunOSS = aliyunOSSConfig{
	Endpoint:        "oss-cn-guangzhou.aliyuncs.com",
	AccessKeyId:     "",
	AccessKeySecret: "",
	BucketName:      "jk-competition",
	BucketUrl:       "https://jk-competition.oss-cn-guangzhou.aliyuncs.com",
	BasePath:        "picture",
}
