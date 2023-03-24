package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

type ChatGptApi struct{}

func (chat *ChatGptApi) CreateToken(c *gin.Context) {

}

func (chat *ChatGptApi) DeleteToken(c *gin.Context) {

}

func (chat *ChatGptApi) GetTable(c *gin.Context) {
	var req request.ChatGptRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	results := chatGptService.GetTable(req)
	response.OkWithDetailed(results, "获取表格成功", c)
}
