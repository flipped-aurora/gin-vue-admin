package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type UserInfoSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
                      UserNickname  string `json:"userNickname" form:"userNickname" `
                      UserGender  *int `json:"userGender" form:"userGender" `
                      UserGrade  *int `json:"userGrade" form:"userGrade" `
                      UserProfession  string `json:"userProfession" form:"userProfession" `
                      UserCity  string `json:"userCity" form:"userCity" `
                      UserModel  *int `json:"userModel" form:"userModel" `
    request.PageInfo
}
