package message

// Video 视频消息
type Video struct {
	CommonToken

	Video struct {
		MediaID     string `xml:"MediaId"`
		Title       string `xml:"Title,omitempty"`
		Description string `xml:"Description,omitempty"`
	} `xml:"Video"`
}

// NewVideo 回复图片消息
func NewVideo(mediaID, title, description string) *Video {
	video := new(Video)
	video.Video.MediaID = mediaID
	video.Video.Title = title
	video.Video.Description = description
	return video
}
