package message

// Voice 语音消息
type Voice struct {
	CommonToken

	Voice struct {
		MediaID string `xml:"MediaId"`
	} `xml:"Voice"`
}

// NewVoice 回复语音消息
func NewVoice(mediaID string) *Voice {
	voice := new(Voice)
	voice.Voice.MediaID = mediaID
	return voice
}
