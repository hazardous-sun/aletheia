package models

type Config struct {
	url     bool
	image   bool
	video   bool
	context bool
}

func NewConfig(url, image, video, context bool) *Config {
	return &Config{url, image, video, context}
}
