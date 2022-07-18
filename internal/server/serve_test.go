package server_test

import (
	"testing"

	"github.com/smlx/go-cli-github/internal/server"
)

func TestServe(t *testing.T) {
	var testCases = map[string]struct {
		input  string
		expect string
	}{
		"test Serve": {input: "", expect: "example serve command"},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			result := server.Serve()
			if result != tc.expect {
				tt.Fatalf("expected %s, got %s", tc.expect, result)
			}
		})
	}
}
