package models_test

import (
	"aletheia-server/src/models"
	"encoding/json"
	"testing"
)

func TestCrawler_EmptyStruct(t *testing.T) {
	crawler := models.Crawler{}
	expected := map[string]interface{}{
		"Id":           float64(0),
		"PagesToVisit": float64(0),
		"Query":        "",
		"QueryUrl":     "",
		"HtmlSelector": "",
		"Status":       "",
		"PagesBodies":  nil, // Empty slice becomes nil in JSON
	}

	testCrawlerMarshaling(t, crawler, expected)
}

func TestCrawler_CompleteStruct(t *testing.T) {
	crawler := models.Crawler{
		Id:           1,
		PagesToVisit: 10,
		Query:        "test query",
		QueryUrl:     "https://example.com/search",
		HtmlSelector: "div.result",
		Status:       "active",
		PagesBodies:  []string{"<html>page1</html>", "<html>page2</html>"},
	}
	expected := map[string]interface{}{
		"Id":           float64(1),
		"PagesToVisit": float64(10),
		"Query":        "test query",
		"QueryUrl":     "https://example.com/search",
		"HtmlSelector": "div.result",
		"Status":       "active",
		"PagesBodies":  []interface{}{"<html>page1</html>", "<html>page2</html>"},
	}

	testCrawlerMarshaling(t, crawler, expected)
}

func TestCrawler_PartialFields(t *testing.T) {
	crawler := models.Crawler{
		Id:          42,
		Query:       "partial test",
		Status:      "pending",
		PagesBodies: []string{"<html>test</html>"},
	}
	expected := map[string]interface{}{
		"Id":           float64(42),
		"PagesToVisit": float64(0),
		"Query":        "partial test",
		"QueryUrl":     "",
		"HtmlSelector": "",
		"Status":       "pending",
		"PagesBodies":  []interface{}{"<html>test</html>"},
	}

	testCrawlerMarshaling(t, crawler, expected)
}

func TestCrawler_JSONFieldNames(t *testing.T) {
	crawler := models.Crawler{
		Id:           1,
		PagesToVisit: 5,
		Query:        "field names test",
	}

	jsonData, err := json.Marshal(crawler)
	if err != nil {
		t.Fatalf("Failed to marshal Crawler to JSON: %v", err)
	}

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(jsonData, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Verify the exact JSON field names are correct (note capitalization)
	expectedFields := []string{"Id", "PagesToVisit", "Query", "QueryUrl", "HtmlSelector", "Status", "PagesBodies"}
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
func testCrawlerMarshaling(t *testing.T, crawler models.Crawler, expected map[string]interface{}) {
	t.Helper()

	jsonData, err := json.Marshal(crawler)
	if err != nil {
		t.Fatalf("Failed to marshal Crawler to JSON: %v", err)
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

		switch v := expectedValue.(type) {
		case []interface{}:
			if v == nil {
				if actualValue != nil {
					t.Errorf("For key '%s': expected nil, got %v", key, actualValue)
				}
			} else {
				if actualValue == nil {
					t.Errorf("For key '%s': expected slice, got nil", key)
				} else {
					if len(v) != len(actualValue.([]interface{})) {
						t.Errorf("For key '%s': expected slice length %d, got %d",
							key, len(v), len(actualValue.([]interface{})))
					}
				}
			}
		default:
			if actualValue != expectedValue {
				t.Errorf("For key '%s': expected %v (%T), got %v (%T)",
					key, expectedValue, expectedValue, actualValue, actualValue)
			}
		}
	}

	// Verify no extra fields
	for key := range unmarshaled {
		if _, ok := expected[key]; !ok {
			t.Errorf("Unexpected JSON key '%s' found in marshaled output", key)
		}
	}
}
