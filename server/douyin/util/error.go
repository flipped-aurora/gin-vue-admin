package util

// CommonError 抖音返回的通用错误.
type CommonError struct {
	ErrCode int64  `json:"error_code"`
	ErrMsg  string `json:"description"`
}

// CommonErrorExtra 抖音返回的错误额外信息.
type CommonErrorExtra struct {
	LogID string `json:"logid"`
	Now   int64  `json:"now"`
}
