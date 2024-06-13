package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/temp"
	"github.com/flipped-aurora/gin-vue-admin/server/service/base"
)

type TempService struct {
	base.BaseService
}

func NewTempService() *TempService {
	return &TempService{
		BaseService: base.BaseService{
			Model: &temp.Temp{},
		},
	}
}

// 其他的自己的业务逻辑函数
func (s *TempService) Q自己定义() string {
	return "OK"
}
