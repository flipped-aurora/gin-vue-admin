package request

import "encoding/json"

type Call struct {
	User    string `json:"user"`
	Soft    string `json:"soft"`
	Command string `json:"command"`

	Data map[string]interface{} `json:"data"`
	Req  string
}

func (c *Call) RequestJSON() (string, error) {
	j, err := json.Marshal(c.Data)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

type ApiCaller struct {
	Call
}
