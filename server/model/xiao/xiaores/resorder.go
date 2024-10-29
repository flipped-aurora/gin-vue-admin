package xiaores

import "github.com/flipped-aurora/gin-vue-admin/server/model/xiao"

type CliOrderRes struct {
	MyMainorder *xiao.CliMainorder `json:"mymainorder" form:"mymainorder"`
	CliChildren []*xiao.CliOrder   `json:"clichildren"`
	CliPulls    []*xiao.CliOrder   `json:"clipulls"`
}

type CliTreeRes struct {
	Cliuser   *xiao.CliUser
	Pulltrees []*xiao.CliTree
	Teamtrees []*xiao.CliTree
}

type CliWithdrawRes struct {
	CliWithdraws []*xiao.CliWithdraw
	CliMainwiths *xiao.CliMainwith
}
