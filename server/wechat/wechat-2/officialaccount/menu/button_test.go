package menu

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewButtonFun(t *testing.T) {
	buttons := []*Button{
		NewSubButton("1", []*Button{
			NewViewButton("1.1", "https://baidu.com"),
			NewViewButton("1.2", "https://baidu.com"),
			NewViewButton("1.3", "https://baidu.com"),
		}),
		NewSubButton("2", []*Button{
			NewViewButton("2.1", "https://baidu.com"),
			NewViewButton("2.2", "https://baidu.com"),
			NewViewButton("2.3", "https://baidu.com"),
		}),
		NewViewButton("3", "https://baidu.com"),
	}

	data, err := json.Marshal(buttons)
	assert.Nil(t, err)
	assert.Equal(t, `[{"name":"1","sub_button":[{"type":"view","name":"1.1","url":"https://baidu.com"},{"type":"view","name":"1.2","url":"https://baidu.com"},{"type":"view","name":"1.3","url":"https://baidu.com"}]},{"name":"2","sub_button":[{"type":"view","name":"2.1","url":"https://baidu.com"},{"type":"view","name":"2.2","url":"https://baidu.com"},{"type":"view","name":"2.3","url":"https://baidu.com"}]},{"type":"view","name":"3","url":"https://baidu.com"}]`, string(data))
}
