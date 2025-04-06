package models_test

import (
	"aletheia-client/src/models"
	"flag"
	"os"
	"testing"
)

func setupEnv(env map[string]string) func() {
	// Backup current environment
	oldEnv := os.Environ()
	os.Clearenv()

	// Set up new environment
	for k, v := range env {
		os.Setenv(k, v)
	}

	// Return cleanup function
	return func() {
		os.Clearenv()
		for _, e := range oldEnv {
			os.Setenv(e, "1")
		}
	}
}

func resetFlags() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
}

func TestNewConfig_WithValidPort(t *testing.T) {
	cleanup := setupEnv(map[string]string{"PORT": "8080"})
	defer cleanup()
	defer resetFlags()

	config, err := models.NewConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if config.Port != "8080" {
		t.Errorf("Expected Port=8080, got %s", config.Port)
	}
}

func TestNewConfig_WithAllFields(t *testing.T) {
	cleanup := setupEnv(map[string]string{"PORT": "8080"})
	defer cleanup()
	defer resetFlags()

	// Set flags
	os.Args = []string{"test", "-P", "-I", "-V"}
	flag.Parse()

	config, err := models.NewConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := models.Config{
		Port:   "8080",
		Prompt: true,
		Image:  true,
		Video:  true,
	}

	if config != expected {
		t.Errorf("Expected %+v, got %+v", expected, config)
	}
}

func TestNewConfig_WithMissingPort(t *testing.T) {
	cleanup := setupEnv(map[string]string{})
	defer cleanup()
	defer resetFlags()

	_, err := models.NewConfig()
	if err == nil {
		t.Error("Expected error for missing PORT, got nil")
	}
}

func TestNewConfig_WithInvalidPort(t *testing.T) {
	cleanup := setupEnv(map[string]string{"PORT": "notanumber"})
	defer cleanup()
	defer resetFlags()

	_, err := models.NewConfig()
	if err == nil {
		t.Error("Expected error for invalid PORT, got nil")
	}
}

func TestNewConfig_WithOptionalFields(t *testing.T) {
	tests := []struct {
		name     string
		env      map[string]string
		args     []string
		expected models.Config
	}{
		{
			"Only PROMPT",
			map[string]string{"PORT": "8080"},
			[]string{"test", "-P"},
			models.Config{Port: "8080", Prompt: true},
		},
		{
			"Only IMAGE",
			map[string]string{"PORT": "8080"},
			[]string{"test", "-I"},
			models.Config{Port: "8080", Image: true},
		},
		{
			"Only VIDEO",
			map[string]string{"PORT": "8080"},
			[]string{"test", "-V"},
			models.Config{Port: "8080", Video: true},
		},
		{
			"PROMPT and IMAGE",
			map[string]string{"PORT": "8080"},
			[]string{"test", "-P", "-I"},
			models.Config{Port: "8080", Prompt: true, Image: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanup := setupEnv(tt.env)
			defer cleanup()
			defer resetFlags()

			// Set flags for this test case
			os.Args = tt.args
			flag.Parse()

			config, err := models.NewConfig()
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if config != tt.expected {
				t.Errorf("Expected %+v, got %+v", tt.expected, config)
			}
		})
	}
}

func TestNewConfig_WithLongFormFlags(t *testing.T) {
	tests := []struct {
		name     string
		env      map[string]string
		args     []string
		expected models.Config
	}{
		{
			"PROMPT long form",
			map[string]string{"PORT": "8080"},
			[]string{"test", "--PROMPT"},
			models.Config{Port: "8080", Prompt: true},
		},
		{
			"IMAGE long form",
			map[string]string{"PORT": "8080"},
			[]string{"test", "--IMAGE"},
			models.Config{Port: "8080", Image: true},
		},
		{
			"VIDEO long form",
			map[string]string{"PORT": "8080"},
			[]string{"test", "--VIDEO"},
			models.Config{Port: "8080", Video: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanup := setupEnv(tt.env)
			defer cleanup()
			defer resetFlags()

			// Set flags for this test case
			os.Args = tt.args
			flag.Parse()

			config, err := models.NewConfig()
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if config != tt.expected {
				t.Errorf("Expected %+v, got %+v", tt.expected, config)
			}
		})
	}
}
