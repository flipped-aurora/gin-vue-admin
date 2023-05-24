package credential

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/douyin/open/cache"
	"github.com/flipped-aurora/gin-vue-admin/server/douyin/util"
)

// 获取ticket的url
const getTicketURL = "https://open.douyin.com/js/getticket?access_token=%s"

// DefaultJsTicket 默认获取js ticket方法
type DefaultJsTicket struct {
	appID          string
	cacheKeyPrefix string
	cache          cache.Cache
	//jsAPITicket 读写锁 同一个AppID一个
	jsAPITicketLock *sync.Mutex
}

// NewDefaultJsTicket new
func NewDefaultJsTicket(appID string, cacheKeyPrefix string, cache cache.Cache) JsTicketHandle {
	return &DefaultJsTicket{
		appID:           appID,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		jsAPITicketLock: new(sync.Mutex),
	}
}

// Ticket 请求jsapi_tikcet返回结果
type Ticket struct {
	util.CommonError

	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// GetTicket 获取jsapi_ticket
func (js *DefaultJsTicket) GetTicket(accessToken string) (ticketStr string, err error) {
	js.jsAPITicketLock.Lock()
	defer js.jsAPITicketLock.Unlock()

	//先从cache中取
	jsAPITicketCacheKey := fmt.Sprintf("%s_jsapi_ticket_%s", js.cacheKeyPrefix, js.appID)
	val := js.cache.Get(jsAPITicketCacheKey)
	if val != nil {
		ticketStr = val.(string)
		return
	}
	var ticket Ticket
	ticket, err = GetTicketFromServer(accessToken)
	if err != nil {
		return
	}
	expires := ticket.ExpiresIn - 1500
	err = js.cache.Set(jsAPITicketCacheKey, ticket.Ticket, time.Duration(expires)*time.Second)
	ticketStr = ticket.Ticket
	return
}

// GetTicketFromServer 从服务器中获取ticket
func GetTicketFromServer(accessToken string) (ticket Ticket, err error) {
	var response []byte
	url := fmt.Sprintf(getTicketURL, accessToken)
	response, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &ticket)
	if err != nil {
		return
	}
	if ticket.ErrCode != 0 {
		err = fmt.Errorf("getTicket Error : errcode=%d , errmsg=%s", ticket.ErrCode, ticket.ErrMsg)
		return
	}
	return
}
