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

func CollectConfigValues(config Config) []string {
	var result []string

	for _, v := range os.Environ() {
		switch v {
		case "CONTEXT=1":
			fmt.Println(v)
			config.Context = os.Getenv(v) == "1"
			result = append(result, "CONTEXT")
		case "IMAGE":
			fmt.Println(v)
			config.Image = os.Getenv(v) == "1"
			result = append(result, v)
		case "VIDEO":
			fmt.Println(v)
			config.Video = os.Getenv(v) == "1"
			result = append(result, v)
		default:
			continue
		}
	}
	fmt.Println(result)

	return result
}
