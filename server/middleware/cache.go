package middleware

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const KEY_PREFIX = "cache_api:"
const KEY_USERID_BIND = "user_bind_id="
const BodyBytesKey = "_gin-gonic/gin/bodybyteskey"

type Cached struct {
	Status   int
	Body     []byte
	Header   http.Header
	ExpireAt time.Duration
}

type Store interface {
	GetInt(string) (int, error)
	GetString(string) (string, error)
	GetStringWithBytes(key string) ([]byte, error)
	Set(string, interface{}, time.Duration) error
	SetNoExpiration(string, interface{}) error
	Delete(string) error
	DeleteAllKeys(string) error
}

type Options struct {
	Expire        time.Duration
	Headers       []string
	DoNotUseAbort bool
}

func (o *Options) init() {
	if o.Headers == nil {
		o.Headers = []string{
			"Authorization",
		}
	}
}

type Cache struct {
	Store
	options Options
	expires map[string]time.Time
}

func (c *Cache) Get(key string) (*Cached, error) {
	if data, err := c.Store.GetStringWithBytes(key); err == nil {
		var cch *Cached
		dec := gob.NewDecoder(bytes.NewBuffer(data))
		if err := dec.Decode(&cch); err == nil {
			return cch, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (c *Cache) Set(key string, cch *Cached) error {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(*cch); err != nil {
		return err
	}
	return c.Store.Set(key, string(b.Bytes()), cch.ExpireAt)
}

func (c *Cache) DeleteAllKeys(key string) error {
	return c.Store.DeleteAllKeys(key)
}

type wrappedWriter struct {
	gin.ResponseWriter
	body bytes.Buffer
}

func (rw *wrappedWriter) Write(body []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(body)
	if err == nil {
		rw.body.Write(body)
	}
	return n, err
}

//authentication.AuthUser() 该中间件切记一定要放在cache.New(true)之前执行
func CacheApi(isBindUser bool) gin.HandlerFunc {
	opts := Options{
		Expire: 0,
	}
	opts.init()

	cache := Cache{
		Store:   global.GVA_REDIS,
		options: opts,
		expires: make(map[string]time.Time),
	}

	return func(c *gin.Context) {
		// only GET method available for caching
		if c.Request.Method != "GET" && c.Request.Method != "POST" {
			c.Next()
			return
		}

		requestUrl := c.Request.URL.RequestURI()
		if isBindUser {
			for _, k := range cache.options.Headers {
				if k == "Authorization" {
					if v, ok := c.Get("claims"); ok {
						claims := v.(*systemReq.CustomClaims)
						userId := strconv.Itoa(int(claims.ID))
						requestUrl = KEY_USERID_BIND + userId + requestUrl
					}
				}
			}
		}
		// 获取get post 参数通用做法
		// 1.get参数，直接使用url就行
		getParam := md5String(requestUrl)
		// 2.post使用body摘要
		body, err := c.GetRawData()
		if err != nil {
			global.GVA_LOG.Error("获取post body 报错", zap.Error(err))
		}
		// 再设置回去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // 关键点

		postParam := md5String(utils.Bytes2str(body))

		key := KEY_PREFIX + requestUrl + "===" + getParam + "---" + postParam
		//fmt.Println(key)
		if cch, _ := cache.Get(key); cch == nil {
			// cache miss
			writer := c.Writer
			rw := wrappedWriter{ResponseWriter: c.Writer}
			c.Writer = &rw
			c.Next()
			c.Writer = writer
			// 判断如果返回错误的数据不进行缓存
			if response, isExit := c.Get("response_code"); isExit {
				if response.(int) == 1 {
					err := cache.Set(key, &Cached{
						Status: rw.Status(),
						Body:   rw.body.Bytes(),
						Header: rw.Header(),
						ExpireAt: func() time.Duration {
							if cache.options.Expire == 0 {
								return 5 * time.Minute
							} else {
								return cache.options.Expire
							}
						}(),
					})
					if err != nil {
						global.GVA_LOG.Error("写入redis缓存失败", zap.Error(err))
					}
				}
			}
		} else {
			// cache found
			start := time.Now()
			c.Writer.WriteHeader(cch.Status)
			for k, val := range cch.Header {
				for _, v := range val {
					if _, ok := c.Writer.Header()[k]; !ok {
						c.Writer.Header().Add(k, v)
					}
				}
			}
			c.Writer.Header().Add("X-Gin-Cache", fmt.Sprintf("%f ms", time.Now().Sub(start).Seconds()*1000))
			if _, err := c.Writer.Write(cch.Body); err == nil {
				if !cache.options.DoNotUseAbort {
					c.Abort()
				}
			}
		}
	}
}

func md5String(url string) string {
	h := md5.New()
	if _, err := io.WriteString(h, url); err != nil {
		global.GVA_LOG.Error("md5 error", zap.Error(err))
	}
	return hex.EncodeToString(h.Sum(nil))
}
// 清除绑定的用户缓存信息
func ClearUserCache(uid uint) error {
	userId := strconv.Itoa(int(uid))
	key := KEY_PREFIX + KEY_USERID_BIND + userId + "*"
	return global.GVA_REDIS.DeleteAllKeys(key)
}
