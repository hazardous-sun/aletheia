package models_test

import (
	"aletheia-client/src/models"
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

func TestNewConfig_WithValidPort(t *testing.T) {
	cleanup := setupEnv(map[string]string{"PORT": "8080"})
	defer cleanup()

	config, err := models.NewConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if config.Port != "8080" {
		t.Errorf("Expected Port=8080, got %s", config.Port)
	}
}

func TestNewConfig_WithAllFields(t *testing.T) {
	cleanup := setupEnv(map[string]string{
		"PORT":    "8080",
		"CONTEXT": "1",
		"IMAGE":   "1",
		"VIDEO":   "1",
	})
	defer cleanup()

	config, err := models.NewConfig()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := models.Config{
		Port:    "8080",
		Context: true,
		Image:   true,
		Video:   true,
	}

	if config != expected {
		t.Errorf("Expected %+v, got %+v", expected, config)
	}
}

func TestNewConfig_WithMissingPort(t *testing.T) {
	cleanup := setupEnv(map[string]string{})
	defer cleanup()

	_, err := models.NewConfig()
	if err == nil {
		t.Error("Expected error for missing PORT, got nil")
	}
}

func TestNewConfig_WithInvalidPort(t *testing.T) {
	cleanup := setupEnv(map[string]string{"PORT": "notanumber"})
	defer cleanup()

	_, err := models.NewConfig()
	if err == nil {
		t.Error("Expected error for invalid PORT, got nil")
	}
}

func TestNewConfig_WithOptionalFields(t *testing.T) {
	tests := []struct {
		name     string
		env      map[string]string
		expected models.Config
	}{
		{
			"Only CONTEXT",
			map[string]string{"PORT": "8080", "CONTEXT": "1"},
			models.Config{Port: "8080", Context: true},
		},
		{
			"Only IMAGE",
			map[string]string{"PORT": "8080", "IMAGE": "1"},
			models.Config{Port: "8080", Image: true},
		},
		{
			"Only VIDEO",
			map[string]string{"PORT": "8080", "VIDEO": "1"},
			models.Config{Port: "8080", Video: true},
		},
		{
			"CONTEXT and IMAGE",
			map[string]string{"PORT": "8080", "CONTEXT": "1", "IMAGE": "1"},
			models.Config{Port: "8080", Context: true, Image: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanup := setupEnv(tt.env)
			defer cleanup()

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
