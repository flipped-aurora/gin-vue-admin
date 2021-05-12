package shared

type GameResponse struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
