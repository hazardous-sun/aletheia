package models_test

import (
	"aletheia-server/src/models"
	"encoding/json"
	"testing"
)

func TestNewsOutlet_EmptyStruct(t *testing.T) {
	outlet := models.NewsOutlet{}
	expected := map[string]interface{}{
		"id":           float64(0),
		"name":         "",
		"queryUrl":     "",
		"htmlSelector": "",
		"language":     "",
		"credibility":  float64(0),
	}

	testNewsOutletMarshaling(t, outlet, expected)
}

func TestNewsOutlet_CompleteStruct(t *testing.T) {
	outlet := models.NewsOutlet{
		Id:           1,
		Name:         "Reuters",
		QueryUrl:     "https://reuters.com/search?q=",
		HtmlSelector: "div.article",
		Language:     "English",
		Credibility:  90,
	}
	expected := map[string]interface{}{
		"id":           float64(1),
		"name":         "Reuters",
		"queryUrl":     "https://reuters.com/search?q=",
		"htmlSelector": "div.article",
		"language":     "English",
		"credibility":  float64(90),
	}

	testNewsOutletMarshaling(t, outlet, expected)
}

func TestNewsOutlet_PartialFields(t *testing.T) {
	outlet := models.NewsOutlet{
		Id:          42,
		Name:        "Partial News",
		Language:    "Spanish",
		Credibility: 75,
	}
	expected := map[string]interface{}{
		"id":           float64(42),
		"name":         "Partial News",
		"queryUrl":     "",
		"htmlSelector": "",
		"language":     "Spanish",
		"credibility":  float64(75),
	}

	testNewsOutletMarshaling(t, outlet, expected)
}

func TestNewsOutlet_MinMaxCredibility(t *testing.T) {
	tests := []struct {
		name         string
		credibility  int
		expectedJson float64
	}{
		{"MinCredibility", 0, float64(0)},
		{"MaxCredibility", 100, float64(100)},
		{"NegativeCredibility", -10, float64(-10)},
		{"AboveMaxCredibility", 150, float64(150)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outlet := models.NewsOutlet{
				Credibility: tt.credibility,
			}
			expected := map[string]interface{}{
				"id":           float64(0),
				"name":         "",
				"queryUrl":     "",
				"htmlSelector": "",
				"language":     "",
				"credibility":  tt.expectedJson,
			}

			testNewsOutletMarshaling(t, outlet, expected)
		})
	}
}

func TestNewsOutlet_JSONFieldNames(t *testing.T) {
	outlet := models.NewsOutlet{
		Id:       1,
		Name:     "Field Names Test",
		QueryUrl: "https://test.com",
	}

	jsonData, err := json.Marshal(outlet)
	if err != nil {
		t.Fatalf("Failed to marshal NewsOutlet to JSON: %v", err)
	}

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	expectedFields := []string{"id", "name", "queryUrl", "htmlSelector", "language", "credibility"}
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
func testNewsOutletMarshaling(t *testing.T, outlet models.NewsOutlet, expected map[string]interface{}) {
	t.Helper()

	jsonData, err := json.Marshal(outlet)
	if err != nil {
		t.Fatalf("Failed to marshal NewsOutlet to JSON: %v", err)
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
