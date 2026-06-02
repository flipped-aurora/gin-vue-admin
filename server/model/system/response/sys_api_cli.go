package response

type ApiCliDetectedParam struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
}

type ApiCliDetectedParams struct {
	Path    []ApiCliDetectedParam `json:"path"`
	Query   []ApiCliDetectedParam `json:"query"`
	Header  []ApiCliDetectedParam `json:"header"`
	HasBody bool                  `json:"hasBody"`
}

type ApiCliPreviewResponse struct {
	ScriptName     string               `json:"scriptName"`
	ScriptContent  string               `json:"scriptContent"`
	Summary        string               `json:"summary"`
	Description    string               `json:"description"`
	Warnings       []string             `json:"warnings"`
	DetectedParams ApiCliDetectedParams `json:"detectedParams"`
}
