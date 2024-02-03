package service

import (
	"app/myRedis"
	"app/service/competition"
	"app/service/discuss"
	"app/service/user"
	"fmt"
	"reflect"
)

type GetInfoById interface {
	GetInfo(id string) (reData interface{}, err error)
}

// 缓存接口解耦升级
func GetUserInfoInRedisOrMySql(inData GetInfoById, id string) (reData interface{}, err error) {
	resverString := ""
	var nowType interface{}
	switch reflect.TypeOf(inData) {
	case reflect.TypeOf(user.UserInfoService{}):
		resverString = "userID:"
		break
	case reflect.TypeOf(discuss.DisInfoService{}):
		resverString = "disID:"
		break
	case reflect.TypeOf(competition.ComInfoService{}):
		resverString = "comID:"
		break
		//case reflect.TypeOf(user.UserInfoSearch{}):
		//	break
		//case reflect.TypeOf(user.UserInfoSearch{}):
		//	break
		//case reflect.TypeOf(user.UserInfoSearch{}):
		//	break
		//case reflect.TypeOf(user.UserInfoSearch{}):
		//	break
	}
	if resverString == "" {
		panic("接口没写好")
	}
	ret, err := myRedis.GetStructVal(&nowType, fmt.Sprintf("%s%s", resverString, id))
	if err != nil {
		data, err := inData.GetInfo(id)
		if err != nil {
			return data, err
		}
		myRedis.SetStructVal(fmt.Sprintf("%s%s", resverString, id), data, 10)
		return data, nil
	} else {
		return ret, nil

	}
}
