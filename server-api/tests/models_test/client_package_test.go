package models_test

import (
	"aletheia-server/src/models"
	"encoding/json"
	"testing"
)

func TestPackageReceived_EmptyStruct(t *testing.T) {
	pkg := models.PackageReceived{}
	expected := map[string]interface{}{
		"url":         "",
		"description": false,
		"image":       false,
		"vide":        false,
		"prompt":      "",
	}

	testPackageMarshaling(t, pkg, expected)
}

func TestPackageReceived_CompleteStruct(t *testing.T) {
	pkg := models.PackageReceived{
		Url:         "https://example.com",
		Description: true,
		Image:       true,
		Video:       true,
		Prompt:      "test prompt",
	}
	expected := map[string]interface{}{
		"url":         "https://example.com",
		"description": true,
		"image":       true,
		"vide":        true,
		"prompt":      "test prompt",
	}

	testPackageMarshaling(t, pkg, expected)
}

func TestPackageReceived_PartialFields(t *testing.T) {
	pkg := models.PackageReceived{
		Url:    "https://partial.com",
		Image:  true,
		Prompt: "partial test",
	}
	expected := map[string]interface{}{
		"url":         "https://partial.com",
		"description": false,
		"image":       true,
		"vide":        false,
		"prompt":      "partial test",
	}

	testPackageMarshaling(t, pkg, expected)
}

func TestPackageReceived_JSONFieldNames(t *testing.T) {
	pkg := models.PackageReceived{
		Url:         "https://example.com",
		Description: true,
		Image:       false,
		Video:       true,
		Prompt:      "field names test",
	}

	jsonData, err := json.Marshal(pkg)
	if err != nil {
		t.Fatalf("Failed to marshal PackageReceived to JSON: %v", err)
	}

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify the exact JSON field names are correct
	expectedFields := []string{"url", "description", "image", "vide", "prompt"}
	for _, field := range expectedFields {
		if _, ok := unmarshaled[field]; !ok {
			t.Errorf("Expected JSON field '%s' not found", field)
		}
	}

	if len(unmarshaled) != len(expectedFields) {
		t.Errorf("Expected exactly %d JSON fields, got %d", len(expectedFields), len(unmarshaled))
	}
}

// Helper function to test marshaling/unmarshaling behavior
func testPackageMarshaling(t *testing.T, pkg models.PackageReceived, expected map[string]interface{}) {
	t.Helper()

	jsonData, err := json.Marshal(pkg)
	if err != nil {
		t.Fatalf("Failed to marshal PackageReceived to JSON: %v", err)
	}

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify JSON fields and values
	for key, expectedValue := range expected {
		actualValue, ok := unmarshaled[key]
		if !ok {
			t.Errorf("Expected JSON key '%s' not found", key)
			continue
		}

		if actualValue != expectedValue {
			t.Errorf("For key '%s': expected %v (%T), got %v (%T)",
				key, expectedValue, expectedValue, actualValue, actualValue)
		}
	}

	// Verify no extra fields
	for key := range unmarshaled {
		if _, ok := expected[key]; !ok {
			t.Errorf("Unexpected JSON key '%s' found in marshaled output", key)
		}
	}
}
