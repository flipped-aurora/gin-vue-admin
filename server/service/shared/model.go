package shared

type GameRequest struct {
	RspCmd int32       `json:"rspCmd" binding:"required"`
	Data   interface{} `json:"data" binding:"required"`
}

type GameResponse struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
