package tools

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// 保存验证码集合 key == token+phone or email  value == code
const TokenCodeYan = "TokenCodeYan"

// 发送手机验证码限制
const SendPhoneCodeLimit = "SendPhoneCodeLimit"

// 发送邮箱验证码限制
const SendEmailCodeLimit = "SendEmailCodeLimit"

// 网络请求限制
const RequestClientIPLimit = "RequestClientIPLimit"

// 同IP开启上下文请求限制
const RequestClientIPContextLimit = "RequestClientIPContextLimit"

// socket客户端在线保持的心跳
const SocketClientHeartbeat = "SocketClientHeartbeat"

// 登录用户token保存
const ClientToken = "LOGIN--"
const AdminToken = "ADMIN_LOGIN--"
const AgentToken = "AGENT_LOGIN--"


// 广播订阅
const (
	UpdateConfig = "UpdateConfig"
	Other        = "Other"
)


type Redis struct {
	Client *redis.Client
}

func (self *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	err := self.Client.Set(context.Background(), key, value, expiration).Err()
	return err
}

// 添加数组数据
func (self *Redis) SetListString(key string, values []string) error {
	err := self.Client.RPush(context.Background(), key, values).Err()
	return err
}

// 读取数组长度
func (self *Redis) GetListLength(key string) (int64, error) {
	return self.Client.LLen(context.Background(), key).Result()
}

// 保留列表区间元素
func (self *Redis) LtrimListString(key string, start, stop int64) error {
	return self.Client.LTrim(context.Background(), key, start, stop).Err()
}

// 删除最老的数据 删除index最小的元素 或者可以理解为删除栈底元素
func (self *Redis) LPopListString(key string) error {
	return self.Client.LPop(context.Background(), key).Err()
}

// 删除最新的数据 删除index最大的元素 或者可以理解为删除栈底元素
func (self *Redis) RPopListString(key string) error {
	return self.Client.RPop(context.Background(), key).Err()
}

// 获取列表所有元素 所有就是从 0 到 -1
func (self *Redis) GetListLRange(key string, start, stop int64) ([]string, error) {
	return self.Client.LRange(context.Background(), key, start, stop).Result()
}

// 获取列表最新一条元素
func (self *Redis) GetListBRPop(key string) ([]string, error) {
	return self.Client.BRPop(context.Background(), time.Second*1, key).Result()
}

// 判断列表中是否存在该元素
func (self *Redis) CheckListHasValue(key string, value string) bool {
	vs, err := self.GetListLRange(key, 0, -1)
	if err != nil {
		return false
	}
	i := -1
	for index, v := range vs {
		if v == value {
			i = index
		}
	}
	return i != -1
}

// 删除列表中的元素
func (self *Redis) DeleteListHasValue(key string, value string) bool {
	err := self.Client.LRem(context.Background(), key, 0, value).Err()
	return err == nil
}

// 添加次数限制
func (self *Redis) SetLimitWithTime(key string, limit int, expiration time.Duration) error {
	if !self.HasKey(key) {
		pipe := self.Client.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, expiration)
		_, err := pipe.Exec(context.Background())
		return err
	} else {
		//times次数
		if times, err := self.GetInt(key); err != nil {
			return err
		} else {
			if times >= limit {
				if t, err := self.Client.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("操作太频繁请稍后再试")
				} else {
					return errors.New("操作太频繁, 请 " + t.String() + " 秒后再试")
				}
			} else {
				return self.Client.Incr(context.Background(), key).Err()
			}
		}
	}
}

// 是否过期该Key 如果没过期自动续期+1 如果单位时间内值大于阈值就报异常
func (self *Redis) AddLimitWithTimeCycle(key string, limit int, expiration time.Duration) error {
	if !self.HasKey(key) {
		pipe := self.Client.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, expiration)
		_, err := pipe.Exec(context.Background())
		return err
	} else {
		//times次数
		if times, err := self.GetInt(key); err != nil {
			return err
		} else {
			if times >= limit {
				if t, err := self.Client.PTTL(context.Background(), key).Result(); err != nil {
					return errors.New("操作太频繁请稍后再试")
				} else {
					return errors.New("操作太频繁, 请 " + t.String() + " 秒后再试")
				}
			} else {
				return self.Client.Incr(context.Background(), key).Err()
			}
		}
	}
}

// 减一
func (self *Redis) DecrLimitWithTimeCycle(key string) error {
	// 加一
	return self.Client.Decr(context.Background(), key).Err()
}

func (self *Redis) SetNoExpiration(key string, value interface{}) error {
	err := self.Set(key, value, 0)
	return err
}

func (self *Redis) HasKey(key string) bool {
	_, err := self.Client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return false
	} else if err != nil {
		return false
	}
	return true
}

func (self *Redis) GetInt(key string) (int, error) {
	return self.Client.Get(context.Background(), key).Int()
}

func (self *Redis) GetInt64(key string) (int64, error) {
	return self.Client.Get(context.Background(), key).Int64()
}

func (self *Redis) GetString(key string) (string, error) {
	return self.Client.Get(context.Background(), key).Result()
}

func (self *Redis) GetStringWithBytes(key string) ([]byte, error) {
	return self.Client.Get(context.Background(), key).Bytes()
}

func (self *Redis) Delete(key string) error {
	return self.Client.Del(context.Background(), key).Err()
}

// 模糊删除key
func (self *Redis) DeleteAllKeys(pattern string) error {
	fmt.Println("模糊删除redis", pattern)
	if result, err := self.Client.Keys(context.Background(), pattern).Result(); err != nil {
		return err
	} else {
		return self.Client.Del(context.Background(), result...).Err()
	}
}

// 模糊查询
func (self *Redis) QueryAllKeys(pattern string) ([]string, error) {
	return self.Client.Keys(context.Background(), pattern).Result()
}

// redis发送消息
func (self *Redis) SendMessage(channel string, message interface{}) error {
	return self.Client.Publish(context.Background(), channel, message).Err()
}
