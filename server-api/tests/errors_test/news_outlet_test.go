package server_errors

import (
	server_errors "aletheia-server/src/errors"
	"testing"
)

func TestNewsOutletErrorConstants(t *testing.T) {
	tests := []struct {
		name     string
		constant string
		want     string
	}{
		{
			name:     "NewsOutletNotFound",
			constant: server_errors.NewsOutletNotFound,
			want:     "news outlet not found inside the database",
		},
		{
			name:     "NewsOutletAlreadyExists",
			constant: server_errors.NewsOutletAlreadyExists,
			want:     "news outlet already exists inside the database",
		},
		{
			name:     "NewsOutletTableMissing",
			constant: server_errors.NewsOutletTableMissing,
			want:     "news outlet table does not exist inside the database",
		},
		{
			name:     "NewsOutletParsingError",
			constant: server_errors.NewsOutletParsingError,
			want:     "news outlet could not be parsed from the database",
		},
		{
			name:     "NewsOutletClosingTableError",
			constant: server_errors.NewsOutletClosingTableError,
			want:     "news outlet table could not be closed properly",
		},
		{
			name:     "NewsOutletNotAdded",
			constant: server_errors.NewsOutletNotAdded,
			want:     "news outlet was not properly added to the database",
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
