package server_errors

import (
	server_errors "fact-checker-server/src/errors"
	"testing"
)

func TestLanguageErrorConstants(t *testing.T) {
	testCases := []struct {
		name     string
		got      string
		expected string
	}{
		{
			name:     "LanguageNotFound",
			got:      server_errors.LanguageNotFound,
			expected: "language not found inside the database",
		},
		{
			name:     "LanguageAlreadyExists",
			got:      server_errors.LanguageAlreadyExists,
			expected: "language already exists inside the database",
		},
		{
			name:     "LanguageTableMissing",
			got:      server_errors.LanguageTableMissing,
			expected: "language table does not exist inside the database",
		},
		{
			name:     "LanguageParsingError",
			got:      server_errors.LanguageParsingError,
			expected: "language row could not be parsed from the database",
		},
		{
			name:     "LanguageClosingTableError",
			got:      server_errors.LanguageClosingTableError,
			expected: "language table could not be closed properly",
		},
		{
			name:     "LanguageNotAdded",
			got:      server_errors.LanguageNotAdded,
			expected: "language was not properly added to the database",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.expected {
				t.Errorf("constant %s has value %q, want %q",
					tc.name, tc.got, tc.expected)
			}
		})
	}
}
