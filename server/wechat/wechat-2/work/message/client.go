// Package message 消息推送,实现企业微信消息推送相关接口：https://developer.work.weixin.qq.com/document/path/90235
package message

import (
	"github.com/flipped-aurora/gin-vue-admin/server/wechat/wechat-2/work/context"
)

// Client 消息推送接口实例
type Client struct {
	*context.Context
}

// NewClient 初始化实例
func NewClient(ctx *context.Context) *Client {
	return &Client{ctx}
}
