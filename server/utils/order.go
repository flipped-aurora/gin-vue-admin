package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GetTimeTick64() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetTimeTick32() int32 {
	return int32(time.Now().Unix())
}

func GetFormatTime(time time.Time) string {
	return time.Format("20060102")
}

// 基础做法 日期20191025时间戳1571987125435+3位随机数
func GenerateOrderNo() string {
	date := GetFormatTime(time.Now())
	r := rand.Intn(1000)
	code := fmt.Sprintf("%s%d%03d", date, GetTimeTick64(), r)
	//fmt.Println(code, " rand ID generate successed!\n")
	return code
}
