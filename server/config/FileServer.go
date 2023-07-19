package config

type FileServer struct {
	MainLink          string `mapstructure:"main-link" json:"main-link" yaml:"main-link"`
	ThumbnailsImgLink string `mapstructure:"thumbnails-img-link" json:"thumbnails-img-link" yaml:"thumbnails-img-link"`
	OriginVideoLink   string `mapstructure:"origin-video-link" json:"origin-video-link" yaml:"origin-video-link"`
}
