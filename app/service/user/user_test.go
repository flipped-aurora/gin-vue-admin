package user

import (
	"fmt"
	"testing"
)

func TestUserInfoService_GetUserInfoByWxOpenid(t *testing.T) {
	fmt.Println("getUser")
	userS := new(UserInfoService)
	user, err := userS.GetUserInfoByWxOpenid("1")
	if err != nil {
		return
	}
	fmt.Println("user = ", user)

}
func TestGetInfoInRedis(t *testing.T) {

}
