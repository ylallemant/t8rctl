package credentials

import (
	"os"
	"testing"
)

func TestNewCredential(t *testing.T) {
	for _, tt := range []struct {
		name       string
		configured bool
		value      string
		wantValue  string
		wantError  bool
	}{
		{name: "unset defaults to developer credentials", wantValue: "dev"},
		{name: "explicit developer credentials", configured: true, value: "dev", wantValue: "dev"},
		{name: "explicit production credentials", configured: true, value: "prod", wantValue: "prod"},
		{name: "invalid configuration", configured: true, value: "invalid", wantValue: "invalid", wantError: true},
		{name: "explicit empty configuration", configured: true, wantError: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("AZURE_TOKEN_CREDENTIALS", tt.value)
			if !tt.configured {
				if err := os.Unsetenv("AZURE_TOKEN_CREDENTIALS"); err != nil {
					t.Fatal(err)
				}
			}

			cred, err := newCredential()
			if (err != nil) != tt.wantError {
				t.Fatalf("newCredential() error = %v, wantError = %v", err, tt.wantError)
			}
			if !tt.wantError && cred == nil {
				t.Fatal("newCredential() returned nil credentials")
			}
			if value := os.Getenv("AZURE_TOKEN_CREDENTIALS"); value != tt.wantValue {
				t.Fatalf("AZURE_TOKEN_CREDENTIALS = %q, want %q", value, tt.wantValue)
			}
		})
	}
}
