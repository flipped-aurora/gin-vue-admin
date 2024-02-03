package myRedis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RD redis.Conn

func init() {
	c, err := redis.Dial("tcp", "123.207.73.185:6379", redis.DialDatabase(1))
	if err != nil {
		panic(err)
		return
	}

	RD = c
}
func Get(key string) (interface{}, error) {
	return RD.Do("get", key)
}
func SetStructVal(key string, value interface{}, expiration time.Duration) (reply interface{}, err error) {
	datas, _ := json.Marshal(value)
	return RD.Do("set", key, datas)
}
func GetStructVal(st interface{}, key string) (interface{}, error) {
	rebytes, err := redis.Bytes(RD.Do("get", key))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rebytes, st)
	if err != nil {
		return nil, err
	}
	return st, nil
}
