package request

type AutoMcpTool struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Params      []struct {
		Name        string `json:"name" form:"name" binding:"required"`
		Description string `json:"description" form:"description" binding:"required"`
		Type        string `json:"type" form:"type" binding:"required"` // string, number, boolean, object, array
		Required    bool   `json:"required" form:"required"`
		Default     string `json:"default" form:"default"`
	} `json:"params" form:"params"`
	Response []struct {
		Type string `json:"type" form:"type" binding:"required"` // text, image
	} `json:"response" form:"response"`
}
