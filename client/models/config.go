package models

type Config struct {
	image   bool
	video   bool
	context bool
}

func NewConfig(image, video, context bool) *Config {
	return &Config{image, video, context}
}
