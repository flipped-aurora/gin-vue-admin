package service

import (
	"app/service/user"
	"fmt"
	"testing"
	"time"
)

func TestGetUserInfoInRedisOrMySql(t *testing.T) {

	var sss user.UserInfoService
	fmt.Println(time.Now())
	for i := 0; i < 1000; i++ {
		sss.GetUserInfo("1")
		fmt.Println(i)
	}
	fmt.Println(time.Now())
}
