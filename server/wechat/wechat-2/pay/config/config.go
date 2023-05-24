package config

// Config .config for pay
type Config struct {
	AppID     string `json:"app_id"`
	MchID     string `json:"mch_id"`
	Key       string `json:"key"`
	NotifyURL string `json:"notify_url"`
}
