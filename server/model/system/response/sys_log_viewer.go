package response

import "time"

type LogDateItem struct {
	Date      string `json:"date"`
	FileCount int    `json:"fileCount"`
}

type LogDateList struct {
	Month string        `json:"month"`
	Dates []LogDateItem `json:"dates"`
}

type LogFileItem struct {
	Path       string    `json:"path"`
	Name       string    `json:"name"`
	Size       int64     `json:"size"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

type LogFileList struct {
	Date  string        `json:"date"`
	Files []LogFileItem `json:"files"`
}

type LogContent struct {
	Date           string    `json:"date"`
	Path           string    `json:"path"`
	Content        string    `json:"content"`
	LineCount      int       `json:"lineCount"`
	NextCursor     int64     `json:"nextCursor"`
	HasMore        bool      `json:"hasMore"`
	LimitedByBytes bool      `json:"limitedByBytes"`
	Size           int64     `json:"size"`
	ModifiedAt     time.Time `json:"modifiedAt"`
}
