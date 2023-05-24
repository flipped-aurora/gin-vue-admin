package message

// Music 音乐消息
type Music struct {
	CommonToken

	Music struct {
		Title        string `xml:"Title"        `
		Description  string `xml:"Description"  `
		MusicURL     string `xml:"MusicUrl"     `
		HQMusicURL   string `xml:"HQMusicUrl"   `
		ThumbMediaID string `xml:"ThumbMediaId"`
	} `xml:"Music"`
}

// NewMusic  回复音乐消息
func NewMusic(title, description, musicURL, hQMusicURL, thumbMediaID string) *Music {
	music := new(Music)
	music.Music.Title = title
	music.Music.Description = description
	music.Music.MusicURL = musicURL
	music.Music.ThumbMediaID = thumbMediaID
	return music
}
