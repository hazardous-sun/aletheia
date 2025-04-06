package models

import (
	"aletheia-client/src/models"
	"encoding/json"
	"testing"
)

// TestPackageSentInitialization tests basic struct initialization
func TestPackageSentInitialization(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		image    bool
		prompt   string
		video    bool
		expected models.PackageSent
	}{
		{
			name:   "All fields populated",
			url:    "https://example.com",
			image:  true,
			prompt: "test prompt",
			video:  false,
			expected: models.PackageSent{
				Url:    "https://example.com",
				Image:  true,
				Prompt: "test prompt",
				Video:  false,
			},
		},
		{
			name:   "Empty fields",
			url:    "",
			image:  false,
			prompt: "",
			video:  false,
			expected: models.PackageSent{
				Url:    "",
				Image:  false,
				Prompt: "",
				Video:  false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := models.PackageSent{
				Url:    tt.url,
				Image:  tt.image,
				Prompt: tt.prompt,
				Video:  tt.video,
			}

			if p != tt.expected {
				t.Errorf("Expected %+v, got %+v", tt.expected, p)
			}
		})
	}
}

// TestPackageSentJSONSerialization tests JSON marshaling/unmarshaling
func TestPackageSentJSONSerialization(t *testing.T) {
	tests := []struct {
		name     string
		input    models.PackageSent
		expected string
	}{
		{
			name: "All fields",
			input: models.PackageSent{
				Url:    "https://test.com",
				Image:  true,
				Prompt: "test prompt",
				Video:  false,
			},
			expected: `{"url":"https://test.com","image":true,"prompt":"test prompt","video":false}`,
		},
		{
			name:     "Empty struct",
			input:    models.PackageSent{},
			expected: `{"url":"","image":false,"prompt":"","video":false}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test marshaling
			bytes, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}

			if string(bytes) != tt.expected {
				t.Errorf("Expected JSON %s, got %s", tt.expected, string(bytes))
			}

			// Test unmarshaling
			var p models.PackageSent
			err = json.Unmarshal(bytes, &p)
			if err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			if p != tt.input {
				t.Errorf("Expected %+v after unmarshal, got %+v", tt.input, p)
			}
		})
	}
}

// TestPackageSentFieldValidation tests field validation logic
func TestPackageSentFieldValidation(t *testing.T) {
	tests := []struct {
		name    string
		p       models.PackageSent
		isValid bool
	}{
		{
			name: "Valid with URL and prompt",
			p: models.PackageSent{
				Url:    "https://valid.com",
				Prompt: "valid prompt",
			},
			isValid: true,
		},
		{
			name: "Invalid empty URL",
			p: models.PackageSent{
				Url:    "",
				Prompt: "valid prompt",
			},
			isValid: false,
		},
		{
			name: "Invalid empty prompt",
			p: models.PackageSent{
				Url:    "https://valid.com",
				Prompt: "",
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This is a placeholder for actual validation logic
			// In a real application, you would implement this method
			isValid := tt.p.Url != "" && tt.p.Prompt != ""

			if isValid != tt.isValid {
				t.Errorf("Expected isValid=%v, got %v", tt.isValid, isValid)
			}
		})
	}
}

// TestPackageSentEdgeCases tests edge cases
func TestPackageSentEdgeCases(t *testing.T) {
	t.Run("Long URL and prompt", func(t *testing.T) {
		longURL := "https://example.com/" + string(make([]byte, 1000))
		longPrompt := string(make([]byte, 10000))

		p := models.PackageSent{
			Url:    longURL,
			Prompt: longPrompt,
		}

		if p.Url != longURL || p.Prompt != longPrompt {
			t.Error("Failed to handle long strings")
		}
	})

	t.Run("Special characters in fields", func(t *testing.T) {
		specialURL := "https://example.com/测试?param=値&other=😊"
		specialPrompt := "prompt with 测试 and 😊"

		p := models.PackageSent{
			Url:    specialURL,
			Prompt: specialPrompt,
		}

		if p.Url != specialURL || p.Prompt != specialPrompt {
			t.Error("Failed to handle special characters")
		}

		// Test JSON roundtrip with special chars
		bytes, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Failed to marshal: %v", err)
		}

		var p2 models.PackageSent
		err = json.Unmarshal(bytes, &p2)
		if err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if p2.Url != specialURL || p2.Prompt != specialPrompt {
			t.Error("Special characters not preserved in JSON roundtrip")
		}
	})
}
