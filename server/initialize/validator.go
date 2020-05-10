package initialize

import "gin-vue-admin/utils"

func init() {
	_ = utils.RegisterRule("PageVerify",
		utils.Rules{
			"Page":     {utils.NotEmpty()},
			"PageSize": {utils.NotEmpty()},
		},
	)
	_ = utils.RegisterRule("IdVerify",
		utils.Rules{
			"Id": {utils.NotEmpty()},
		},
	)
	_ = utils.RegisterRule("AuthorityIdVerify",
		utils.Rules{
			"AuthorityId": {utils.NotEmpty()},
		},
	)
}
