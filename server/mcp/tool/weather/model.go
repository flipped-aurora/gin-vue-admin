package weather

type Tool struct{}

type Request struct {
	City string `json:"city" description:"city name"`
	Date string `json:"date" description:"date"`
}
