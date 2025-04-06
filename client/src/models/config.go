package models

import (
	"aletheia-client/src/errors"
	"errors"
	"flag"
	"fmt"
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

	// Get the PORT value from the environment variables
	err := getPort(&config)

	if err != nil {
		client_errors.Log(err.Error(), client_errors.ErrorLevel)
		return Config{}, err
	}

	// Define the flags
	contextFlag := flag.Bool("C", false, "Context parameter")
	contextFlagLong := flag.Bool("CONTEXT", false, "Context parameter (long form)")
	imageFlag := flag.Bool("I", false, "Image parameter")
	imageFlagLong := flag.Bool("IMAGE", false, "Image parameter (long form)")
	videoFlag := flag.Bool("V", false, "Video parameter")
	videoFlagLong := flag.Bool("VIDEO", false, "Video parameter (long form)")

	// Parse the flags
	flag.Parse()

	// Check which flags were set
	if *contextFlag || *contextFlagLong {
		config.Context = true
	}

	if *imageFlag || *imageFlagLong {
		config.Image = true
	}

	if *videoFlag || *videoFlagLong {
		config.Video = true
	}

	// Check if the Port field was initialized
	if config.Port == "" {
		err := errors.New(client_errors.UninitializedPort)
		client_errors.Log(client_errors.UninitializedPort, client_errors.WarningLevel)
		return Config{}, err
	}

	warnMissingFields(config)

	return config, nil
}

func getPort(config *Config) error {
	value, _ := os.LookupEnv(Port)
	// Log the port value being used
	client_errors.Log(fmt.Sprintf("PORT = '%s'", value), client_errors.InfoLevel)

	// Pass value to the struct
	config.Port = value

	// Check if the port is a valid integer
	_, err := strconv.Atoi(config.Port)

	if err != nil {
		err := errors.New(client_errors.InvalidPortValue)
		return err
	}

	config.Port = value
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
