package google_secrets

import (
	"testing"
)

func TestGetSecret(t *testing.T) {
	secret, err := get("core-359619", "google-secret-integration-test")
	expectedString := "All right, then.  Keep your secrets."
	if err != nil {
		t.Error(err)
	}
	if secret != expectedString {
		t.Errorf("Expected String(%s) is not the same as actual string (%s).", secret, expectedString)
	}
}
