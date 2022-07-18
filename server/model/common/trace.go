package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

type Trace struct {
	TraceId   string  `json:"trace_id"`
	SpanId    string  `json:"span_id"`
	Caller    string  `json:"caller"`
	SrcMethod *string `json:"srcMethod,omitempty"`
	UserId    int     `json:"user_id"`
}

var TraceCtx = "traceCtx"
var TraceKey = "traceKey"

func GetTraceCtx(c *gin.Context) context.Context {
	if value, ok := c.Get(TraceCtx); ok {
		return value.(context.Context)
	} else {
		uuidStr := strings.ReplaceAll(uuid.New().String(), "-", "")
		return context.WithValue(context.Background(), TraceKey, &Trace{TraceId: uuidStr})
	}
}

//	获取traceId
func GetTraceId(c *gin.Context) string {
	var uuidStr string
	if c != nil {
		if value, ok := c.Get(TraceCtx); ok {
			trace := value.(context.Context).Value(TraceKey).(*Trace)
			uuidStr = trace.TraceId
		} else {
			uuidStr = strings.ReplaceAll(uuid.New().String(), "-", "")
			context.WithValue(context.Background(), TraceKey, &Trace{TraceId: uuidStr})
		}
	}

	return uuidStr
}
