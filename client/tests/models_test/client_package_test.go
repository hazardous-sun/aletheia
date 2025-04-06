package models_test

import (
	"aletheia-client/src/models"
	"encoding/json"
	"testing"
)

// TestPackageSent_Empty verifies zero values
func TestPackageSent_Empty(t *testing.T) {
	p := models.PackageSent{}

	if p.Url != "" {
		t.Errorf("Expected empty Url, got '%s'", p.Url)
	}
	if p.Image {
		t.Error("Expected false Image, got true")
	}
	if p.Video {
		t.Error("Expected false Video, got true")
	}
	if p.Prompt != "" {
		t.Errorf("Expected empty Prompt, got '%s'", p.Prompt)
	}
}

// TestPackageSent_FieldAssignment verifies field assignments
func TestPackageSent_FieldAssignment(t *testing.T) {
	p := models.PackageSent{
		Url:    "https://test.com",
		Image:  false,
		Video:  true,
		Prompt: "test",
	}

	if p.Url != "https://test.com" {
		t.Errorf("Expected Url 'https://test.com', got '%s'", p.Url)
	}
	if p.Image {
		t.Error("Expected Image false, got true")
	}
	if !p.Video {
		t.Error("Expected Video true, got false")
	}
	if p.Prompt != "test" {
		t.Errorf("Expected Prompt 'test', got '%s'", p.Prompt)
	}
}

// TestPackageSent_JSONMarshal tests JSON marshaling
func TestPackageSent_JSONMarshal(t *testing.T) {
	p := models.PackageSent{
		Url:    "https://marshal.com",
		Video:  true,
		Prompt: "marshal test",
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	expected := `{"url":"https://marshal.com","image":false,"vide":true,"prompt":"marshal test"}`
	if string(data) != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, string(data))
	}
}

// TestPackageSent_JSONUnmarshal tests JSON unmarshaling
func TestPackageSent_JSONUnmarshal(t *testing.T) {
	jsonStr := `{
		"url": "https://unmarshal.com",
		"vide": false,
		"prompt": "unmarshal test"
	}`

	var p models.PackageSent
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if p.Url != "https://unmarshal.com" {
		t.Errorf("Expected Url 'https://unmarshal.com', got '%s'", p.Url)
	}
	if p.Video {
		t.Error("Expected Video false, got true")
	}
	if p.Prompt != "unmarshal test" {
		t.Errorf("Expected Prompt 'unmarshal test', got '%s'", p.Prompt)
	}
}

// TestPackageSent_JSONTagTypo verifies the intentional "vide" JSON tag
func TestPackageSent_JSONTagTypo(t *testing.T) {
	p := models.PackageSent{Video: true}
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if _, exists := result["vide"]; !exists {
		t.Error("Expected 'vide' field in JSON output")
	}
	if val, exists := result["vide"]; !exists || !val.(bool) {
		t.Error("Expected 'vide' field to be true")
	}
}
