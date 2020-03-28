package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"gin-vue-admin/init/initlog"
)

type Config struct {
	MysqlAdmin   MysqlAdmin   `json:"mysqlAdmin"`
	Qiniu        Qiniu        `json:"qiniu"`
	CasbinConfig CasbinConfig `json:"casbinConfig"`
	RedisAdmin   RedisAdmin   `json:"redisAdmin"`
	System       System       `json:"system"`
	JWT          JWT          `json:"jwt"`
	Captcha      Captcha      `json:"captcha"`
	Log          Log          `json:"log"`
}

type System struct { // 系统配置
	UseMultipoint bool   `json:"useMultipoint"`
	Env           string `json:"env"`
	Addr          int    `json:"addr"`
}

type JWT struct { // jwt签名
	SigningKey string `json:"signingKey"`
}

type CasbinConfig struct { //casbin配置
	ModelPath string `json:"modelPath"` // casbin model地址配置
}

type MysqlAdmin struct { // mysql admin 数据库配置
	Username     string `json:"username"`
	Password     string `json:"password"`
	Path         string `json:"path"`
	Dbname       string `json:"dbname"`
	Config       string `json:"config"`
	MaxIdleConns int    `json:"maxIdleConns"`
	MaxOpenConns int    `json:"maxOpenConns"`
	LogMode      bool   `json:"maxOpenConns"`
	Retry        Retry
}

type Retry struct {
	Count int
	Wait  time.Duration
}

type RedisAdmin struct { // Redis admin 数据库配置
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
type Qiniu struct { // 七牛 密钥配置
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type Captcha struct { // 验证码配置
	KeyLong   int `json:"keyLong"`
	ImgWidth  int `json:"imgWidth"`
	ImgHeight int `json:"imgHeight"`
}

/**
Log Config

"CRITICAL"
"ERROR"
"WARNING"
"NOTICE"
"INFO"
"DEBUG"
*/
type Log struct {
	// log 打印的前缀
	Prefix string `json:"prefix"`
	// 是否显示打印log的文件具体路径
	LogFile bool `json:"logFile"`
	// 在控制台打印log的级别， []默认不打印
	Stdout []string `json:"stdout"`
	// 在文件中打印log的级别   []默认不打印
	File []string `json:"file"`
}

var GinVueAdminconfig Config
var VTool *viper.Viper

func init() {
	v := viper.New()
	v.SetConfigName("config")           //  设置配置文件名 (不带后缀)
	v.AddConfigPath("./static/config/") // 第一个搜索路径
	v.SetConfigType("json")
	err := v.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&GinVueAdminconfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&GinVueAdminconfig); err != nil {
		fmt.Println(err)
	}
	VTool = v
}

// SetPath 设置MySQL数据库的URL，默认设置为“localhost:3306”。
func (config *MysqlAdmin) SetPath(host, port string) {

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3306"
	}
	config.Path = host + ":" + port
}

// SetRetry 设置重试的次数和间隔，默认设置为不进行重试。
func (config *MysqlAdmin) SetRetry(count, delay string) {

	if count == "" || delay == "" {
		config.Retry.Count = 0
		config.Retry.Wait = 0 * time.Second
		return
	}

	cnt, err := strconv.Atoi(count)
	if err != nil {
		log.L.Fatal(err.Error())
	}
	config.Retry.Count = cnt

	wait, err := strconv.Atoi(delay)
	if err != nil {
		log.L.Fatal(err.Error())
	}
	config.Retry.Wait = time.Duration(wait) * time.Second
}
