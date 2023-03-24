package system

type ChatGpt struct {
	DBName string `json:"dbname,omitempty"`
	Chat   string `json:"chat,omitempty"`
	ChatID string `json:"chatID,omitempty"`
}

type ChatField struct {
	TABLE_NAME     string
	COLUMN_NAME    string
	COLUMN_COMMENT string
}

type ChatFieldNoTable struct {
	COLUMN_NAME    string
	COLUMN_COMMENT string
}
