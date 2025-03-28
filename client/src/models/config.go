package models

import (
	"os"
)

type Config struct {
	Image   bool
	Video   bool
	Context bool
}

func NewConfig(image, video, context bool) *Config {
	return &Config{image, video, context}
}

func CollectConfigValues(config Config) []string {
	var result []string

	for _, v := range os.Environ() {
		switch v {
		case "--CONTEXT":
			config.Context = os.Getenv(v) == "1"
			result = append(result, v)
		case "--IMAGE":
			config.Image = os.Getenv(v) == "1"
			result = append(result, v)
		case "--VIDEO":
			config.Video = os.Getenv(v) == "1"
			result = append(result, v)
		default:
			continue
		}
	}

	return result
}
