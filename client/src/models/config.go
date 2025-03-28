package models

import (
	"fmt"
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

func CollectConfigValues() ([]string, Config) {
	var result []string
	config := Config{false, false, false}

	for _, v := range os.Environ() {
		switch v {
		case "CONTEXT=1":
			config.Context = true
			result = append(result, "CONTEXT")
		case "IMAGE=1":
			config.Image = true
			result = append(result, "IMAGE")
		case "VIDEO=1":
			fmt.Println(v)
			config.Video = true
			result = append(result, "VIDEO")
		default:
			continue
		}
	}

	return result, config
}
