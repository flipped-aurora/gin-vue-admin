package request

type InitDB struct {
	SqlType  string `json:"sqlType"`                     //数据库类型
	Path     string `json:"path"`                        //sqlite数据库路径
	Host     string `json:"host"`                        // 服务器地址
	Port     string `json:"port"`                        // 数据库连接端口
	UserName string `json:"userName" binding:"required"` // 数据库用户名
	Password string `json:"password"`                    // 数据库密码
	DBName   string `json:"dbName" binding:"required"`   // 数据库名
}
