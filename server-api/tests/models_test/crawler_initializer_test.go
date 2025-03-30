package models_test

import (
	"encoding/json"
	"fact-checker-server/src/models"
	"testing"
)

func TestCrawlerInitializer_EmptyStruct(t *testing.T) {
	initializer := models.CrawlerInitializer{}
	expected := map[string]interface{}{
		"query":        "",
		"pagesToVisit": float64(0),
	}

	testCrawlerInitializerMarshaling(t, initializer, expected)
}

func TestCrawlerInitializer_CompleteStruct(t *testing.T) {
	initializer := models.CrawlerInitializer{
		Query:        "test query",
		PagesToVisit: 5,
	}
	expected := map[string]interface{}{
		"query":        "test query",
		"pagesToVisit": float64(5),
	}

	testCrawlerInitializerMarshaling(t, initializer, expected)
}

func TestCrawlerInitializer_OnlyQuery(t *testing.T) {
	initializer := models.CrawlerInitializer{
		Query: "only query test",
	}
	expected := map[string]interface{}{
		"query":        "only query test",
		"pagesToVisit": float64(0),
	}

	testCrawlerInitializerMarshaling(t, initializer, expected)
}

func TestCrawlerInitializer_OnlyPagesToVisit(t *testing.T) {
	initializer := models.CrawlerInitializer{
		PagesToVisit: 10,
	}
	expected := map[string]interface{}{
		"query":        "",
		"pagesToVisit": float64(10),
	}

	testCrawlerInitializerMarshaling(t, initializer, expected)
}

func TestCrawlerInitializer_JSONFieldNames(t *testing.T) {
	initializer := models.CrawlerInitializer{
		Query:        "field names test",
		PagesToVisit: 3,
	}

	jsonData, err := json.Marshal(initializer)
	if err != nil {
		t.Fatalf("Failed to marshal CrawlerInitializer to JSON: %v", err)
	}

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify the exact JSON field names are correct
	expectedFields := []string{"query", "pagesToVisit"}
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
func testCrawlerInitializerMarshaling(t *testing.T, initializer models.CrawlerInitializer, expected map[string]interface{}) {
	t.Helper()

	jsonData, err := json.Marshal(initializer)
	if err != nil {
		t.Fatalf("Failed to marshal CrawlerInitializer to JSON: %v", err)
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
