package modelInterface

// 因为我也不确定项目要不要多人维护 所以定义了CURD接口 凡是对数据库进行简单CURD操作 请实现此接口 默认首位返回 error
type CURD interface {
	Create() (error, CURD)
	Updata() (error, CURD)
	Read() (error, CURD)
	Delete() (error, CURD)
}
