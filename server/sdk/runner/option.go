package runner

func NewConfig(opts ...Option) *Config {
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

type Config struct {
	IsPublicApi bool                   `json:"is_public_api"`
	ApiDesc     string                 `json:"api_desc"`
	MockData    map[string]interface{} `json:"mock_data"`
}
type Option func(*Config)

func WithIsPublicApi(isPublic bool) Option {
	return func(c *Config) {
		c.IsPublicApi = isPublic
	}
}
func WithApiDesc(apiDesc string) Option {
	return func(c *Config) {
		c.ApiDesc = apiDesc
	}
}
func WithMockData(mockData map[string]interface{}) Option {
	return func(c *Config) {
		c.MockData = mockData
	}
}
