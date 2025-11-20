package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-key-123")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if key != "test-key-123" {
		t.Errorf("Expected 'test-key-123', got: %s", key)
	}
}
