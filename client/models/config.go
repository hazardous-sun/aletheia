package models

type Config struct {
	Image   bool
	Video   bool
	Context bool
}

func NewConfig(image, video, context bool) *Config {
	return &Config{image, video, context}
}
