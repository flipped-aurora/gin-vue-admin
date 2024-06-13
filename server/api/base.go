package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Base struct {
	Service interface{}
	Model   interface{}
}

// Create 创建记录
func (b *Base) Create(c *gin.Context) {
	model := b.Model
	err := c.ShouldBindJSON(&model)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	model.(interface {
		SetCreatedBy(userID uint)
	}).SetCreatedBy(utils.GetUserID(c))

	if err := b.Service.(interface {
		Create(interface{}) error
	}).Create(model); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Delete 删除记录
func (b *Base) Delete(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := b.Service.(interface {
		Delete(string, uint) error
	}).Delete(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteByIds 批量删除记录
func (b *Base) DeleteByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := b.Service.(interface {
		DeleteByIds([]string, uint) error
	}).DeleteByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// Update 更新记录
func (b *Base) Update(c *gin.Context) {
	model := b.Model
	err := c.ShouldBindJSON(&model)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	model.(interface {
		SetUpdatedBy(userID uint)
	}).SetUpdatedBy(utils.GetUserID(c))

	if err := b.Service.(interface {
		Update(interface{}) error
	}).Update(model); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// Find 查询记录
func (b *Base) Find(c *gin.Context) {
	ID := c.Query("ID")
	if record, err := b.Service.(interface {
		Find(string) (interface{}, error)
	}).Find(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"record": record}, c)
	}
}

// List 分页查询记录
func (b *Base) List(c *gin.Context, pageInfo interface{}) {
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := b.Service.(interface {
		List(interface{}) ([]interface{}, int64, error)
	}).List(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
			Page: pageInfo.(interface {
				GetPage() int
			}).GetPage(),
			PageSize: pageInfo.(interface {
				GetPageSize() int
			}).GetPageSize(),
		}, "获取成功", c)
	}
}
