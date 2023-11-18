package global

import "time"

type IBlackCache interface {
	Set(k string, v interface{}, d time.Duration)
	SetDefault(k string, v interface{})
	Get(k string) (interface{}, bool)

	Increment(k string, n int64) error
}
