package request

type LogMonthQuery struct {
	Month string `form:"month" binding:"required"`
}

type LogDateQuery struct {
	Date string `form:"date" binding:"required"`
}

type LogContentQuery struct {
	Date   string `form:"date" binding:"required"`
	Path   string `form:"path" binding:"required"`
	Cursor *int64 `form:"cursor"`
}
