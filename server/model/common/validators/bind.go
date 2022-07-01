package validators

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func GVAShouldBind(s interface{}, c *gin.Context) bool {
	if err := c.ShouldBind(s); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回s
			global.GVA_LOG.Error("参数校验报错:", zap.Error(errs))
			response.FailWithMessage("参数校验错误", c)
			c.Abort()
			return true
		}
		// validator.ValidationErrors类型错误则进行翻译
		response.FailWithMessage(RemoveTopStruct(errs.Translate(Trans)), c)
		c.Abort()
		return true
	}
	return false
}

func firstError(fields map[string]string) string {  // 是这个意义？
	//res := map[string]string{}
	//for field, err := range fields {
	//	res[field[strings.Index(field, ".")+1:]] = err
	//	break
	//}
	name := "参数错误"
	for _, err := range fields {
		name = err
		break
	}
	return name
}
