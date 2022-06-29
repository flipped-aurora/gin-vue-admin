//go:build embed_resources
// +build embed_resources

// 仅当使用 embed_resources tag 时，才打包相应的文件

package main

import (
	"embed"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var (
	//go:embed resource
	resource embed.FS

	//go:embed config.yaml config.docker.yaml
	config embed.FS
)

func init() {
	// 如果想延迟 embed FS 的回写，参考：
	// https://stackoverflow.com/questions/66285635/how-do-you-use-go-1-16-embed-features-in-subfolders-packages
	utils.NewEmbedFSBatch(config, resource).RestoreAll()
}
