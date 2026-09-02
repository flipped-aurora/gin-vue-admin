package response

import "time"

type SysAIWorkflowSessionListItem struct {
	ID             uint      `json:"ID"`
	CreatedAt      time.Time `json:"CreatedAt"`
	UpdatedAt      time.Time `json:"UpdatedAt"`
	Tab            string    `json:"tab"`
	Title          string    `json:"title"`
	Summary        string    `json:"summary"`
	ConversationID string    `json:"conversationId"`
	CurrentNodeID  string    `json:"currentNodeId"`
}

type AIWorkflowMarkdownDumpResult struct {
	FileName     string `json:"fileName"`
	FilePath     string `json:"filePath"`
	RelativePath string `json:"relativePath"`
	Directory    string `json:"directory"`
}
