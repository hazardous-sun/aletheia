package client_errors_test

import (
	"aletheia-client/src/errors"
	"bytes"
	"log"
	"testing"
)

// TestLogLevels verifies all log level outputs
func TestLogLevels(t *testing.T) {
	tests := []struct {
		name     string
		message  string
		level    string
		expected string
	}{
		{
			name:     "Info level",
			message:  "test info message",
			level:    client_errors.InfoLevel,
			expected: "\033[0;36minfo: test info message \033[0m",
		},
		{
			name:     "Warning level",
			message:  "test warning message",
			level:    client_errors.WarningLevel,
			expected: "\033[93mwarning: test warning message \033[0m",
		},
		{
			name:     "Error level",
			message:  "test error message",
			level:    client_errors.ErrorLevel,
			expected: "\033[91merror: test error message \033[0m",
		},
	}

	// Redirect log output
	old := log.Writer()
	defer log.SetOutput(old)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log.SetOutput(&buf)

			client_errors.Log(tt.message, tt.level)

			// Remove timestamp from log output
			output := buf.String()[20:] // Skip timestamp (len "2009/11/10 23:00:00 ")

			if output != tt.expected+"\n" {
				t.Errorf("Expected %q, got %q", tt.expected, output)
			}
		})
	}
}

// TestConstants verifies constant values
func TestConstants(t *testing.T) {
	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{"InfoLevel", client_errors.InfoLevel, "info"},
		{"WarningLevel", client_errors.WarningLevel, "warning"},
		{"ErrorLevel", client_errors.ErrorLevel, "error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.actual != tt.expected {
				t.Errorf("Expected %s to be %q, got %q", tt.name, tt.expected, tt.actual)
			}
		})
	}
}
