package models

import (
	client_errors "fact-ckert-client/src/errors"
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
		case "CONTEXT":
			config.Context = os.Getenv("CONTEXT") == "1"
			result = append(result, v)
		case "IMAGE":
			config.Image = os.Getenv("IMAGE") == "1"
			result = append(result, v)
		case "VIDEO":
			config.Video = os.Getenv("VIDEO") == "1"
			result = append(result, v)
		default:
			client_errors.Log(fmt.Sprintf("%s -> %s", client_errors.InvalidEnvVariable, v), client_errors.ErrorLevel)
			panic(client_errors.InvalidEnvVariable)
		}
	}

	return result
}
