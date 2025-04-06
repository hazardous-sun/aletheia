package models

import (
	"aletheia-server/src/models"
	"encoding/json"
	"testing"
)

// Testmodels.PackageReceivedInitialization tests basic struct initialization
func TestPackageReceivedInitialization(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		image    bool
		video    bool
		prompt   string
		expected models.PackageReceived
	}{
		{
			name:   "All fields populated",
			url:    "https://api.example.com/data",
			image:  true,
			video:  false,
			prompt: "generate a landscape image",
			expected: models.PackageReceived{
				Url:    "https://api.example.com/data",
				Image:  true,
				Video:  false,
				Prompt: "generate a landscape image",
			},
		},
		{
			name:   "Video only",
			url:    "https://api.example.com/video",
			image:  false,
			video:  true,
			prompt: "create a short clip",
			expected: models.PackageReceived{
				Url:    "https://api.example.com/video",
				Image:  false,
				Video:  true,
				Prompt: "create a short clip",
			},
		},
		{
			name:   "Empty fields",
			url:    "",
			image:  false,
			video:  false,
			prompt: "",
			expected: models.PackageReceived{
				Url:    "",
				Image:  false,
				Video:  false,
				Prompt: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := models.PackageReceived{
				Url:    tt.url,
				Image:  tt.image,
				Video:  tt.video,
				Prompt: tt.prompt,
			}

			if p != tt.expected {
				t.Errorf("Expected %+v, got %+v", tt.expected, p)
			}
		})
	}
}

// Testmodels.PackageReceivedJSONSerialization tests JSON marshaling/unmarshaling
func TestPackageReceivedJSONSerialization(t *testing.T) {
	tests := []struct {
		name     string
		input    models.PackageReceived
		expected string
	}{
		{
			name: "All fields",
			input: models.PackageReceived{
				Url:    "https://api.test.com",
				Image:  true,
				Video:  false,
				Prompt: "test serialization",
			},
			expected: `{"url":"https://api.test.com","image":true,"video":false,"prompt":"test serialization"}`,
		},
		{
			name: "Image and video false",
			input: models.PackageReceived{
				Url:    "https://api.test.com",
				Image:  false,
				Video:  false,
				Prompt: "no media",
			},
			expected: `{"url":"https://api.test.com","image":false,"video":false,"prompt":"no media"}`,
		},
		{
			name:     "Empty struct",
			input:    models.PackageReceived{},
			expected: `{"url":"","image":false,"video":false,"prompt":""}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test marshaling
			bytes, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}

			// Unmarshal both the expected and actual JSON to compare the structures
			var expected, actual models.PackageReceived
			if err := json.Unmarshal([]byte(tt.expected), &expected); err != nil {
				t.Fatalf("Failed to unmarshal expected JSON: %v", err)
			}
			if err := json.Unmarshal(bytes, &actual); err != nil {
				t.Fatalf("Failed to unmarshal actual JSON: %v", err)
			}

			if actual != expected {
				t.Errorf("Expected %+v, got %+v", expected, actual)
			}

			// Test unmarshaling roundtrip
			var p models.PackageReceived
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

// TestPackageReceivedFieldValidation tests field validation logic
func TestPackageReceivedFieldValidation(t *testing.T) {
	tests := []struct {
		name    string
		p       models.PackageReceived
		isValid bool
	}{
		{
			name: "Valid with URL and prompt",
			p: models.PackageReceived{
				Url:    "https://valid.com",
				Prompt: "valid prompt",
			},
			isValid: true,
		},
		{
			name: "Invalid empty URL",
			p: models.PackageReceived{
				Url:    "",
				Prompt: "valid prompt",
			},
			isValid: false,
		},
		{
			name: "Invalid empty prompt",
			p: models.PackageReceived{
				Url:    "https://valid.com",
				Prompt: "",
			},
			isValid: false,
		},
		{
			name: "Invalid image and video both true",
			p: models.PackageReceived{
				Url:    "https://valid.com",
				Image:  true,
				Video:  true,
				Prompt: "conflicting media types",
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Placeholder for actual validation logic
			// Implement your specific validation requirements here
			isValid := tt.p.Url != "" &&
				tt.p.Prompt != "" &&
				!(tt.p.Image && tt.p.Video) // Example: can't have both image and video true

			if isValid != tt.isValid {
				t.Errorf("Expected isValid=%v, got %v", tt.isValid, isValid)
			}
		})
	}
}

// TestPackageReceivedEdgeCases tests edge cases
func TestPackageReceivedEdgeCases(t *testing.T) {
	t.Run("Maximum length fields", func(t *testing.T) {
		longURL := "https://example.com/" + string(make([]byte, 2000))
		longPrompt := string(make([]byte, 10000))

		p := models.PackageReceived{
			Url:    longURL,
			Prompt: longPrompt,
		}

		if p.Url != longURL || p.Prompt != longPrompt {
			t.Error("Failed to handle maximum length strings")
		}
	})

	t.Run("Unicode and special characters", func(t *testing.T) {
		specialURL := "https://example.com/路径/文件?参数=值&other=😀"
		specialPrompt := "prompt with 汉字, emoji 😀 and symbols &%$"

		p := models.PackageReceived{
			Url:    specialURL,
			Prompt: specialPrompt,
		}

		if p.Url != specialURL || p.Prompt != specialPrompt {
			t.Error("Failed to handle Unicode and special characters")
		}

		// Test JSON roundtrip with special chars
		bytes, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Failed to marshal: %v", err)
		}

		var p2 models.PackageReceived
		err = json.Unmarshal(bytes, &p2)
		if err != nil {
			t.Fatalf("Failed to unmarshal: %v", err)
		}

		if p2.Url != specialURL || p2.Prompt != specialPrompt {
			t.Error("Unicode characters not preserved in JSON roundtrip")
		}
	})

	t.Run("Boolean edge cases", func(t *testing.T) {
		p1 := models.PackageReceived{Image: true, Video: true}
		p2 := models.PackageReceived{Image: false, Video: false}
		p3 := models.PackageReceived{Image: true, Video: false}
		p4 := models.PackageReceived{Image: false, Video: true}

		if !p1.Image || !p1.Video {
			t.Error("Failed to set both booleans true")
		}
		if p2.Image || p2.Video {
			t.Error("Failed to set both booleans false")
		}
		if !p3.Image || p3.Video {
			t.Error("Failed to set image true and video false")
		}
		if p4.Image || !p4.Video {
			t.Error("Failed to set image false and video true")
		}
	})
}
