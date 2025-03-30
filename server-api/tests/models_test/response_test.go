package models_test

import (
	"encoding/json"
	"fact-checker-server/src/models"
	"testing"
)

func TestResponse_EmptyStruct(t *testing.T) {
	resp := models.Response{}
	expected := map[string]interface{}{
		"message": "",
		"status":  float64(0),
	}

	testResponseMarshaling(t, resp, expected)
}

func TestResponse_CompleteStruct(t *testing.T) {
	resp := models.Response{
		Message: "Operation successful",
		Status:  200,
	}
	expected := map[string]interface{}{
		"message": "Operation successful",
		"status":  float64(200),
	}

	testResponseMarshaling(t, resp, expected)
}

func TestResponse_OnlyMessage(t *testing.T) {
	resp := models.Response{
		Message: "Not found",
	}
	expected := map[string]interface{}{
		"message": "Not found",
		"status":  float64(0),
	}

	testResponseMarshaling(t, resp, expected)
}

func TestResponse_OnlyStatus(t *testing.T) {
	resp := models.Response{
		Status: 404,
	}
	expected := map[string]interface{}{
		"message": "",
		"status":  float64(404),
	}

	testResponseMarshaling(t, resp, expected)
}

func TestResponse_StatusCodes(t *testing.T) {
	tests := []struct {
		name         string
		status       int
		expectedJson float64
	}{
		{"Success", 200, float64(200)},
		{"BadRequest", 400, float64(400)},
		{"NotFound", 404, float64(404)},
		{"ServerError", 500, float64(500)},
		{"NegativeStatus", -1, float64(-1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := models.Response{
				Status: tt.status,
			}
			expected := map[string]interface{}{
				"message": "",
				"status":  tt.expectedJson,
			}

			testResponseMarshaling(t, resp, expected)
		})
	}
}

func TestResponse_JSONFieldNames(t *testing.T) {
	resp := models.Response{
		Message: "Test field names",
		Status:  201,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Failed to marshal Response to JSON: %v", err)
	}

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	expectedFields := []string{"message", "status"}
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
func testResponseMarshaling(t *testing.T, resp models.Response, expected map[string]interface{}) {
	t.Helper()

	jsonData, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Failed to marshal Response to JSON: %v", err)
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
