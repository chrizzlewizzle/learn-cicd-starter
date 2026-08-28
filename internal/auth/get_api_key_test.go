package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-token")
	apikey, err := GetAPIKey(headers)

	if apikey != "my-secret-token" {
		t.Fatalf("expected %s, got %s", "my-secret-token", apikey)
	}

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
