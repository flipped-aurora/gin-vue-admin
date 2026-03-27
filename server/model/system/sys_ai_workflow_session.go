package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common"
)

type AIWorkflowMessage struct {
	ID             string         `json:"id"`
	Role           string         `json:"role"`
	Content        string         `json:"content"`
	Snapshot       common.JSONMap `json:"snapshot"`
	ConversationID string         `json:"conversationId"`
	MessageID      string         `json:"messageId"`
	CreatedAt      string         `json:"createdAt"`
}

type SysAIWorkflowSession struct {
	global.GVA_MODEL
	UserID         uint                `json:"userId" gorm:"column:user_id;index;comment:用户ID"`
	Tab            string              `json:"tab" gorm:"column:tab;size:32;index;comment:会话类型"`
	Title          string              `json:"title" gorm:"column:title;size:255;comment:会话标题"`
	Summary        string              `json:"summary" gorm:"column:summary;type:text;comment:摘要"`
	ConversationID string              `json:"conversationId" gorm:"column:conversation_id;size:255;comment:Dify会话ID"`
	MessageID      string              `json:"messageId" gorm:"column:message_id;size:255;comment:Dify消息ID"`
	CurrentNodeID  string              `json:"currentNodeId" gorm:"column:current_node_id;size:64;comment:当前选中节点ID"`
	Settings       common.JSONMap      `json:"settings" gorm:"column:settings;type:longtext;comment:页面设置"`
	FormData       common.JSONMap      `json:"formData" gorm:"column:form_data;type:longtext;comment:表单数据"`
	ResultData     common.JSONMap      `json:"resultData" gorm:"column:result_data;type:longtext;comment:当前展示结果"`
	Messages       []AIWorkflowMessage `json:"messages" gorm:"column:messages;serializer:json;type:longtext;comment:会话消息"`
}

func (s *SysAIWorkflowSession) TableName() string {
	return "sys_ai_workflow_sessions"
}
