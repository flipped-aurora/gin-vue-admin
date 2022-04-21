package model

type AliModel struct {
	Phones        []string `json:"phones"`
	TemplateCode  string   `json:"templateCode"`
	TemplateParam string   `json:"templateParam"`
}
