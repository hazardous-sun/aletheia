package server_errors

import (
	"aletheia-server/src/errors"
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestLoggingConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		want     string
	}{
		{
			name:     "InfoLevel",
			constant: server_errors.InfoLevel,
			want:     "info",
		},
		{
			name:     "WarningLevel",
			constant: server_errors.WarningLevel,
			want:     "warning",
		},
		{
			name:     "ErrorLevel",
			constant: server_errors.ErrorLevel,
			want:     "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.want {
				t.Errorf("%s = %q, want %q", tt.name, tt.constant, tt.want)
			}
		})
	}
}

func TestLogFunction(t *testing.T) {
	// Redirect log output to buffer
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	tests := []struct {
		name     string
		message  string
		level    string
		expected string
	}{
		{
			name:     "InfoLog",
			message:  "test info message",
			level:    server_errors.InfoLevel,
			expected: "info: test info message",
		},
		{
			name:     "WarningLog",
			message:  "test warning message",
			level:    server_errors.WarningLevel,
			expected: "warning: test warning message",
		},
		{
			name:     "ErrorLog",
			message:  "test error message",
			level:    server_errors.ErrorLevel,
			expected: "error: test error message",
		},
		{
			name:     "UnknownLevel",
			message:  "test unknown message",
			level:    "unknown",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset() // Clear buffer before each test
			server_errors.Log(tt.message, tt.level)

			if tt.level == "unknown" {
				if buf.Len() != 0 {
					t.Error("Expected no output for unknown log level")
				}
				return
			}

			output := buf.String()
			if !strings.Contains(output, tt.expected) {
				t.Errorf("Log() output = %q, want it to contain %q", output, tt.expected)
			}
		})
	}
}
