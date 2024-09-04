package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CallResponse struct {
	StatusCode  int         `json:"status_code"`
	Msg         string      `json:"msg"`
	ContentType string      `json:"content_type"`
	HasFile     bool        `json:"has_file"`
	FilePath    string      `json:"path"`
	DeleteFile  bool        `json:"delete_file"`
	Data        interface{} `json:"data"`
}
