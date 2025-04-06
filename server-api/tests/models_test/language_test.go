package models_test

import (
	"aletheia-server/src/models"
	"encoding/json"
	"testing"
)

func TestLanguage_EmptyStruct(t *testing.T) {
	lang := models.Language{}
	expected := map[string]interface{}{
		"id":   float64(0),
		"name": "",
	}

	testLanguageMarshaling(t, lang, expected)
}

func TestLanguage_CompleteStruct(t *testing.T) {
	lang := models.Language{
		Id:   1,
		Name: "Go",
	}
	expected := map[string]interface{}{
		"id":   float64(1),
		"name": "Go",
	}

	testLanguageMarshaling(t, lang, expected)
}

func TestLanguage_OnlyID(t *testing.T) {
	lang := models.Language{
		Id: 42,
	}
	expected := map[string]interface{}{
		"id":   float64(42),
		"name": "",
	}

	testLanguageMarshaling(t, lang, expected)
}

func TestLanguage_OnlyName(t *testing.T) {
	lang := models.Language{
		Name: "Python",
	}
	expected := map[string]interface{}{
		"id":   float64(0),
		"name": "Python",
	}

	testLanguageMarshaling(t, lang, expected)
}

func TestLanguage_JSONFieldNames(t *testing.T) {
	lang := models.Language{
		Id:   1,
		Name: "Go",
	}

	jsonData, err := json.Marshal(lang)
	if err != nil {
		t.Fatalf("Failed to marshal Language to JSON: %v", err)
	}

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify the exact JSON field names are correct
	if _, ok := unmarshaled["id"]; !ok {
		t.Error("Expected JSON field 'id' not found")
	}
	if _, ok := unmarshaled["name"]; !ok {
		t.Error("Expected JSON field 'name' not found")
	}
	if len(unmarshaled) != 2 {
		t.Errorf("Expected exactly 2 JSON fields, got %d", len(unmarshaled))
	}
}

// Helper function to test marshaling/unmarshaling behavior
func testLanguageMarshaling(t *testing.T, lang models.Language, expected map[string]interface{}) {
	t.Helper()

	jsonData, err := json.Marshal(lang)
	if err != nil {
		t.Fatalf("Failed to marshal Language to JSON: %v", err)
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
