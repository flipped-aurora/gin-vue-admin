package system

import (
	"context"
	"errors"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	system "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	systemResp "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"gorm.io/gorm"
)

type aiWorkflowSession struct{}

var AIWorkflowSessionServiceApp = new(aiWorkflowSession)

func (s *aiWorkflowSession) Save(ctx context.Context, userID uint, info systemReq.SysAIWorkflowSessionUpsert) (session system.SysAIWorkflowSession, err error) {
	if userID == 0 {
		return session, errors.New("用户未登录")
	}
	if info.Tab != "analysis" && info.Tab != "workflow" {
		return session, errors.New("不支持的会话类型")
	}

	db := global.GVA_DB.WithContext(ctx)
	if info.ID != 0 {
		err = db.Where("id = ? AND user_id = ?", info.ID, userID).First(&session).Error
		if err != nil {
			return session, err
		}
	}

	session.UserID = userID
	session.Tab = info.Tab
	session.Title = truncateText(firstNonEmpty(strings.TrimSpace(info.Title), s.titleFromMessages(info.Messages), s.titleFromForm(info)), 255)
	session.Summary = strings.TrimSpace(firstNonEmpty(info.Summary, s.summaryFromResult(info.ResultData)))
	session.ConversationID = strings.TrimSpace(info.ConversationID)
	session.MessageID = strings.TrimSpace(info.MessageID)
	session.Settings = cloneJSONMap(info.Settings)
	session.FormData = cloneJSONMap(info.FormData)
	session.ResultData = cloneJSONMap(info.ResultData)
	session.Messages = sanitizeMessages(info.Messages)
	session.CurrentNodeID = strings.TrimSpace(info.CurrentNodeID)
	if session.CurrentNodeID == "" {
		session.CurrentNodeID = lastAssistantMessageID(session.Messages)
	}

	if session.ID == 0 {
		err = db.Create(&session).Error
		return session, err
	}
	err = db.Save(&session).Error
	return session, err
}

func (s *aiWorkflowSession) GetList(ctx context.Context, userID uint, info systemReq.SysAIWorkflowSessionSearch) (list []systemResp.SysAIWorkflowSessionListItem, total int64, err error) {
	db := global.GVA_DB.WithContext(ctx).Model(&system.SysAIWorkflowSession{}).Where("user_id = ?", userID)
	if tab := strings.TrimSpace(info.Tab); tab != "" {
		db = db.Where("tab = ?", tab)
	}
	if keyword := strings.TrimSpace(info.Keyword); keyword != "" {
		like := "%" + keyword + "%"
		db = db.Where("title LIKE ? OR summary LIKE ?", like, like)
	}

	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Select("id", "created_at", "updated_at", "tab", "title", "summary", "conversation_id", "current_node_id").
		Scopes(info.Paginate()).
		Order("updated_at desc").
		Find(&list).Error
	return list, total, err
}

func (s *aiWorkflowSession) GetDetail(ctx context.Context, userID uint, id uint) (session system.SysAIWorkflowSession, err error) {
	err = global.GVA_DB.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).First(&session).Error
	return session, err
}

func (s *aiWorkflowSession) Delete(ctx context.Context, userID uint, id uint) error {
	result := global.GVA_DB.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Delete(&system.SysAIWorkflowSession{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *aiWorkflowSession) titleFromMessages(messages []system.AIWorkflowMessage) string {
	for _, item := range messages {
		if item.Role == "user" && strings.TrimSpace(item.Content) != "" {
			return strings.TrimSpace(item.Content)
		}
	}
	return ""
}

func (s *aiWorkflowSession) titleFromForm(info systemReq.SysAIWorkflowSessionUpsert) string {
	if info.Tab == "analysis" {
		if requirement, ok := info.FormData["requirement"].(string); ok {
			return strings.TrimSpace(requirement)
		}
	}
	if source, ok := info.FormData["source"].(string); ok {
		return strings.TrimSpace(source)
	}
	return ""
}

func (s *aiWorkflowSession) summaryFromResult(resultData map[string]interface{}) string {
	if resultData == nil {
		return ""
	}
	if summary, ok := resultData["summary"].(string); ok {
		return strings.TrimSpace(summary)
	}
	return ""
}

func sanitizeMessages(messages []system.AIWorkflowMessage) []system.AIWorkflowMessage {
	if len(messages) == 0 {
		return []system.AIWorkflowMessage{}
	}
	result := make([]system.AIWorkflowMessage, 0, len(messages))
	for _, item := range messages {
		result = append(result, system.AIWorkflowMessage{
			ID:             strings.TrimSpace(item.ID),
			Role:           strings.TrimSpace(item.Role),
			Content:        item.Content,
			Snapshot:       cloneJSONMap(item.Snapshot),
			ConversationID: strings.TrimSpace(item.ConversationID),
			MessageID:      strings.TrimSpace(item.MessageID),
			CreatedAt:      strings.TrimSpace(item.CreatedAt),
		})
	}
	return result
}

func cloneJSONMap(source map[string]interface{}) map[string]interface{} {
	if len(source) == 0 {
		return map[string]interface{}{}
	}
	target := make(map[string]interface{}, len(source))
	for key, value := range source {
		target[key] = value
	}
	return target
}

func lastAssistantMessageID(messages []system.AIWorkflowMessage) string {
	for i := len(messages) - 1; i >= 0; i-- {
		if messages[i].Role == "assistant" && strings.TrimSpace(messages[i].ID) != "" {
			return messages[i].ID
		}
	}
	return ""
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func truncateText(value string, size int) string {
	if size <= 0 {
		return ""
	}
	runes := []rune(strings.TrimSpace(value))
	if len(runes) <= size {
		return string(runes)
	}
	return string(runes[:size])
}
