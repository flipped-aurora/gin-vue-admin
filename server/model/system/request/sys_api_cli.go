package request

type GenerateApiCliRequest struct {
	Path           string `json:"path"`
	Method         string `json:"method"`
	WrapperName    string `json:"wrapperName"`
	BaseUrlEnvName string `json:"baseUrlEnvName"`
	TokenEnvName   string `json:"tokenEnvName"`
}
