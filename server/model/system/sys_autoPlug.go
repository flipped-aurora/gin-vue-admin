package system

type AutoPlug struct {
	PlugName    string `json:"plugName"`
	RouterGroup string `json:"routerGroup"`
	HasGlobal   bool   `json:"hasGlobal"`
	HasRequest  bool   `json:"hasRequest"`
	HasResponse bool   `json:"hasResponse"`
	Global      []struct {
		Key  string `json:"key"`
		Type string `json:"type"`
	} `json:"global"`
	Request []struct {
		Key  string `json:"key"`
		Type string `json:"type"`
	} `json:"request"`
	Response []struct {
		Key  string `json:"key"`
		Type string `json:"type"`
	} `json:"response"`
}
