package main

type CliConfig struct {
	BaseURL      string `json:"baseURL"`
	Token        string `json:"token"`
	AuthHeader   string `json:"authHeader"`
	ManifestPath string `json:"manifestPath"`
}

type Manifest struct {
	Name     string            `json:"name"`
	Version  string            `json:"version"`
	Server   ManifestServer    `json:"server"`
	Commands []ManifestCommand `json:"commands"`
}

type ManifestServer struct {
	BaseURL    string `json:"baseURL"`
	AuthHeader string `json:"authHeader"`
}

type ManifestCommand struct {
	Name        string              `json:"name"`
	Summary     string              `json:"summary"`
	Description string              `json:"description"`
	Method      string              `json:"method"`
	Path        string              `json:"path"`
	ContentType string              `json:"contentType"`
	Parameters  []ManifestParameter `json:"parameters"`
}

type ManifestParameter struct {
	Name        string `json:"name"`
	Flag        string `json:"flag"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Field       string `json:"field"`
}
