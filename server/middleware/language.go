package middleware

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
)

// Language handler
func LanguageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.Request.Header.Get("Accept-Language")

		// fmt.Printf("Accept-Language value is: %s\r\n", lang)
		var oldLang = global.GVA_CONFIG.Language
		if lang != "" && (lang == "en" || lang == "zh" || lang == "ar") { // currently we support en, zh, ar languages only
			global.GVA_CONFIG.Language = lang
		} else {
			global.GVA_CONFIG.Language = "en"
		}

		if oldLang != lang {
			fmt.Printf("Changing language from: %s to %s\r\n", oldLang, lang)
		}

		global.GVA_TRANSLATOR.SetTranslatorLanguage(global.GVA_CONFIG.Language)

		c.Next()
	}
}
