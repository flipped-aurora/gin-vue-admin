package temp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/temp"
	"github.com/flipped-aurora/gin-vue-admin/server/service/user"
	"github.com/gin-gonic/gin"
)

type Temp struct {
	api.Base
}

type TempApi struct {
}

// 注册路由的时候使用
func NewTemapApi() *Temp {
	return &Temp{
		Base: api.Base{
			Service: user.NewTempService(),
			Model:   &temp.Temp{},
		},
	}
}

// Temp 创建Temp表
// @Tags Temp
// @Summary 创建Temp表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body temp.Temp true "创建Temp表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /ttttt/ttttt [post]
func (temp *Temp) CreateTemp(c *gin.Context) {
	temp.Create(c)
}

func (temp *Temp) FindFunc(c *gin.Context) {
	temp.Find(c)
}
func (temp *Temp) Q自己定义api(c *gin.Context) {
	// 调用服务层的自定义业务逻辑函数
	result := temp.Service.(*user.TempService).Q自己定义()
	response.OkWithMessage(result, c)
}
