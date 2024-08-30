package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/sdk/runner"
	"os"
	"path"
)

var okMap = map[string]string{
	"msg": "ok",
}

func main() {
	r := runner.New()
	r.Post("/file_create", func(ctx *runner.Context) {
		ctx.Logger().Infof("file_create")
		reqMap := ctx.ReqMap()
		filePath := reqMap["file_path"].(string)
		fileName := reqMap["file_name"].(string)
		fileContent := reqMap["file_content"].(string)
		err := os.MkdirAll(filePath, 0755)
		if err != nil {
			ctx.Logger().Error(err)
			ctx.ResponseFailDefaultJSONWithMsg(err.Error())
			return
		}
		filepath := path.Join(filePath, fileName)
		file, err := os.Create(filepath)
		if err != nil {
			ctx.Logger().Error(err)
			ctx.ResponseFailDefaultJSONWithMsg(err.Error())
			return
		}
		defer file.Close()
		_, err = file.WriteString(fileContent)
		if err != nil {
			ctx.Logger().Error(err)
			ctx.ResponseFailDefaultJSONWithMsg(err.Error())
			return
		}
		ctx.Logger().Infof("file_create success, filepath: %s", filepath)
		ctx.ResponseOkWithJSON(okMap)
	})

	r.Post("/file_delete", func(ctx *runner.Context) {
		ctx.Logger().Infof("file_delete")
		reqMap := ctx.ReqMap()
		filePath := reqMap["file"].(string)
		err := os.Remove(filePath)
		if err != nil {
			ctx.Logger().Error(err)
			ctx.ResponseFailDefaultJSONWithMsg(err.Error())
			return
		}

		ctx.Logger().Infof("file_delete success, filepath: %s", filePath)
		ctx.ResponseOkWithJSON(okMap)
	})

	r.Get("/file_get", func(ctx *runner.Context) {
		ctx.Logger().Infof("file_get")
		reqMap := ctx.ReqMap()
		filePath := reqMap["file"].(string)
		file, err := os.ReadFile(filePath)
		//err := os.Remove(filePath)
		if err != nil {
			ctx.Logger().Error(err)
			ctx.ResponseFailDefaultJSONWithMsg(err.Error())
			return
		}

		ctx.Logger().Infof("file_get success, filepath: %s", filePath)
		ctx.ResponseOkWithJSON(map[string]string{
			"file_content": string(file),
		})
	})
}
