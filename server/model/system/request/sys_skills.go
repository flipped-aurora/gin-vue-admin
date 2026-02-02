package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type SkillToolRequest struct {
	Tool string `json:"tool"`
}

type SkillDetailRequest struct {
	Tool  string `json:"tool"`
	Skill string `json:"skill"`
}

type SkillSaveRequest struct {
	Tool      string           `json:"tool"`
	Skill     string           `json:"skill"`
	Meta      system.SkillMeta `json:"meta"`
	Markdown  string           `json:"markdown"`
	SyncTools []string         `json:"syncTools"`
}

type SkillScriptCreateRequest struct {
	Tool       string `json:"tool"`
	Skill      string `json:"skill"`
	FileName   string `json:"fileName"`
	ScriptType string `json:"scriptType"`
}

type SkillResourceCreateRequest struct {
	Tool     string `json:"tool"`
	Skill    string `json:"skill"`
	FileName string `json:"fileName"`
}

type SkillReferenceCreateRequest struct {
	Tool     string `json:"tool"`
	Skill    string `json:"skill"`
	FileName string `json:"fileName"`
}

type SkillTemplateCreateRequest struct {
	Tool     string `json:"tool"`
	Skill    string `json:"skill"`
	FileName string `json:"fileName"`
}

type SkillFileRequest struct {
	Tool     string `json:"tool"`
	Skill    string `json:"skill"`
	FileName string `json:"fileName"`
}

type SkillFileSaveRequest struct {
	Tool     string `json:"tool"`
	Skill    string `json:"skill"`
	FileName string `json:"fileName"`
	Content  string `json:"content"`
}

type SkillGlobalConstraintSaveRequest struct {
	Tool      string   `json:"tool"`
	Content   string   `json:"content"`
	SyncTools []string `json:"syncTools"`
}
