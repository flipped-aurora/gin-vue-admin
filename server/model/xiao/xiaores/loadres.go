package xiaores

import "github.com/flipped-aurora/gin-vue-admin/server/model/xiao"

type CliLoadResponse struct {
	Load *xiao.CliLoad `json:"load" form:"load"`
	User *xiao.CliUser `json:"user" form:"user"`
}
