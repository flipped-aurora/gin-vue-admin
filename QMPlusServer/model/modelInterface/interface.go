package modelInterface

// 因为我也不确定项目要不要多人维护 所以定义了CURD接口 作为接口参考
// 由于很多接口使用Restful模式 暂时不用泛型 有需要可以iss提供示例

type PageInfo struct {
	Page     int
	PageSize int
}

//分页接口
type Paging interface {
	GetInfoList(PageInfo) (err error, list interface{}, total int)
}
