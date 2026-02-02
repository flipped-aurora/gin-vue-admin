package system

type SkillMeta struct {
	Name         string `json:"name" yaml:"name"`
	Description  string `json:"description" yaml:"description"`
	AllowedTools string `json:"allowedTools" yaml:"allowed-tools,omitempty"`
	Context      string `json:"context" yaml:"context,omitempty"`
	Agent        string `json:"agent" yaml:"agent,omitempty"`
}

type SkillDetail struct {
	Tool       string    `json:"tool"`
	Skill      string    `json:"skill"`
	Meta       SkillMeta `json:"meta"`
	Markdown   string    `json:"markdown"`
	Scripts    []string  `json:"scripts"`
	Resources  []string  `json:"resources"`
	References []string  `json:"references"`
	Templates  []string  `json:"templates"`
}

type SkillTool struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}
