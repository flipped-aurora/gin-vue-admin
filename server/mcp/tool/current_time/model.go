package current_time

type Tool struct{}

type Req struct {
	Timezone string `json:"timezone" description:"current time timezone"`
}
