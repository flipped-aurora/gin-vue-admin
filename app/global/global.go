package global

// mysql配置
const DSN = "jk-competition:njhFLS3Bbe8AJZEs@tcp(123.207.73.185:3306)/jk-competition?charset=utf8mb4&parseTime=True&loc=Local"

// JWT配置
const JWTKey = "ab_project"

// 小程序配置
const APPID = "wxa554745366849234"
const SECRET = "3217677b2f40b5e4cb08f1117a9ded36"

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
	AccessKeyId:     "LTAI5tPV5EHihghq4QMu6NRY",
	AccessKeySecret: "49dUJAvg5P9yCGrlIBpUtXVo8c6Rcx",
	BucketName:      "jk-competition",
	BucketUrl:       "https://jk-competition.oss-cn-guangzhou.aliyuncs.com",
	BasePath:        "picture",
}
