package models_test

import (
	"fact-ckert-client/src/models"
	"fmt"
	"os"
	"testing"
)

// TestNewConfig verifies the Config constructor
func TestNewConfig(t *testing.T) {
	tests := []struct {
		name     string
		image    bool
		video    bool
		context  bool
		expected models.Config
	}{
		{"All true", true, true, true, models.Config{true, true, true}},
		{"All false", false, false, false, models.Config{false, false, false}},
		{"Mixed", true, false, true, models.Config{true, false, true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := models.NewConfig(tt.image, tt.video, tt.context)
			if *config != tt.expected {
				t.Errorf("Expected %+v, got %+v", tt.expected, *config)
			}
		})
	}
}

// TestCollectConfigValues_EmptyEnv tests with empty environment
func TestCollectConfigValues_EmptyEnv(t *testing.T) {
	// Backup and clear environment
	oldEnv := os.Environ()
	os.Clearenv()
	defer func() {
		os.Clearenv()
		for _, e := range oldEnv {
			os.Setenv(e, "1")
		}
	}()

	envVars, config := models.CollectConfigValues()

	if len(envVars) != 0 {
		t.Errorf("Expected empty envVars, got %v", envVars)
	}
	if config != (models.Config{false, false, false}) {
		t.Errorf("Expected zero config, got %+v", config)
	}
}

// TestCollectConfigValues_SingleVar tests with single environment variable
func TestCollectConfigValues_SingleVar(t *testing.T) {
	// Backup and modify environment
	oldEnv := os.Environ()
	os.Clearenv()
	os.Setenv("IMAGE", "1")
	defer func() {
		os.Clearenv()
		for _, e := range oldEnv {
			os.Setenv(e, "1")
		}
	}()

	envVars, config := models.CollectConfigValues()

	if len(envVars) != 1 || envVars[0] != "IMAGE" {
		t.Errorf("Expected [IMAGE], got %v", envVars)
	}
	if config != (models.Config{true, false, false}) {
		t.Errorf("Expected {Image:true}, got %+v", config)
	}
}

// TestCollectConfigValues_AllVars tests with all supported environment variables
func TestCollectConfigValues_AllVars(t *testing.T) {
	// Backup and modify environment
	oldEnv := os.Environ()
	os.Clearenv()
	os.Setenv("IMAGE", "1")
	os.Setenv("VIDEO", "1")
	os.Setenv("CONTEXT", "1")
	defer func() {
		os.Clearenv()
		for _, e := range oldEnv {
			os.Setenv(e, "1")
		}
	}()

	envVars, config := models.CollectConfigValues()

	if len(envVars) != 3 {
		t.Errorf("Expected 3 env vars, got %d", len(envVars))
	}
	expectedConfig := models.Config{true, true, true}
	if config != expectedConfig {
		t.Errorf("Expected %+v, got %+v", expectedConfig, config)
	}
}

// TestCollectConfigValues_UnrelatedVars tests with unrelated environment variables
func TestCollectConfigValues_UnrelatedVars(t *testing.T) {
	// Backup and modify environment
	oldEnv := os.Environ()
	os.Clearenv()
	os.Setenv("PATH", "/bin")
	os.Setenv("HOME", "/user")
	defer func() {
		os.Clearenv()
		for _, e := range oldEnv {
			os.Setenv(e, "1")
		}
	}()

	envVars, config := models.CollectConfigValues()

	if len(envVars) != 0 {
		t.Errorf("Expected no env vars, got %v", envVars)
	}
	if config != (models.Config{false, false, false}) {
		t.Errorf("Expected zero config, got %+v", config)
	}
}

// TestCollectConfigValues_VideoOutput verifies VIDEO=1 prints to stdout
func TestCollectConfigValues_VideoOutput(t *testing.T) {
	// Backup and modify environment
	oldEnv := os.Environ()
	os.Clearenv()
	os.Setenv("VIDEO", "1")
	defer func() {
		os.Clearenv()
		for _, e := range oldEnv {
			os.Setenv(e, "1")
		}
	}()

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	models.CollectConfigValues()

	w.Close()
	var output string
	fmt.Fscanf(r, "%s", &output)

	expected := "VIDEO=1"
	if output != expected {
		t.Errorf("Expected output '%s', got '%s'", expected, output)
	}
}
