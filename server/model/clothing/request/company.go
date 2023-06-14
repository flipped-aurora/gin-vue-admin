package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/clothing"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CompanySearch struct {
	clothing.Company
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type JoinCompany struct {
	RoleID    uint   `json:"roleID" form:"roleID"`
	Remark    string `json:"remark" form:"remark"`
	CompanyID uint   `json:"companyID" form:"companyID"`
	TeamID    uint   `json:"teamID" form:"teamID"`
}

type DeleteStaff struct {
	CompanyID uint `json:"companyID" form:"companyID"`
	UserID    uint `json:"userID" form:"userID"`
}
