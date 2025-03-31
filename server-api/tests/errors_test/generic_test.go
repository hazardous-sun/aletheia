package server_errors

import (
	server_errors "fact-checker-server/src/errors"
	"testing"
)

func TestServerErrorConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		want     string
	}{
		{
			name:     "EmptyIdError",
			constant: server_errors.EmptyIdError,
			want:     "id cannot be empty",
		},
		{
			name:     "InvalidIdError",
			constant: server_errors.InvalidIdError,
			want:     "id should be an integer",
		},
		{
			name:     "EmptyNameError",
			constant: server_errors.EmptyNameError,
			want:     "name cannot be empty",
		},
		{
			name:     "InvalidParameters",
			constant: server_errors.InvalidParameters,
			want:     "invalid parameters",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.want {
				t.Errorf("%s: got %q, want %q", tt.name, tt.constant, tt.want)
			}
		})
	}
}
