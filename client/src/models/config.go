package models

import (
	"errors"
	"fact-ckert-client/src/errors"
	"os"
	"strconv"
)

type Config struct {
	Port    string `json:"port"`
	Image   bool   `json:"image"`
	Video   bool   `json:"video"`
	Context bool   `json:"context"`
}

const (
	Port    = "PORT"
	Context = "CONTEXT"
	Image   = "IMAGE"
	Video   = "VIDEO"
)

// NewConfig :
// Returns an instance of a Config struct, used to configure the GUI for the client application.
// Will fail if the "PORT" environment variable is not initialized as a valid integer.
func NewConfig() (Config, error) {
	config := Config{"", false, false, false}
	// Get the values for each field from the environment variables
	err := getValue(&config, Port)

	if err != nil {
		client_errors.Log(err.Error(), client_errors.ErrorLevel)
		return Config{}, err
	}

	_ = getValue(&config, Context)
	_ = getValue(&config, Image)
	_ = getValue(&config, Video)

	// Check if the Port field was initialized
	if config.Port == "" {
		err := errors.New(client_errors.UninitializedPort)
		client_errors.Log(client_errors.UninitializedPort, client_errors.WarningLevel)
		return Config{}, err
	}

	warnMissingFields(config)

	return config, nil
}

func getValue(config *Config, field string) error {
	value, exists := os.LookupEnv(field)
	switch field {
	case Port:
		// Check if the port is a valid integer
		_, err := strconv.Atoi(config.Port)

		if err != nil {
			err := errors.New(client_errors.InvalidPortValue)
			return err
		}

		config.Port = value
	case Context:
		config.Context = exists
	case Image:
		config.Image = exists
	case Video:
		config.Video = exists
	}

	return nil
}

// Warn if there are fields that will not be displayed
func warnMissingFields(config Config) {
	if !config.Context {
		client_errors.Log(client_errors.UninitializedContext, client_errors.WarningLevel)
	}
	if !config.Image {
		client_errors.Log(client_errors.UninitializedImage, client_errors.WarningLevel)
	}
	if !config.Video {
		client_errors.Log(client_errors.UninitializedVideo, client_errors.WarningLevel)
	}
}
