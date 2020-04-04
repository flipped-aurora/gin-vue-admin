package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	code int
	data interface{}
	msg  string
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, response{
		code,
		data,
		msg,
	})
}
