package server_test

import (
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"github.com/smlx/go-cli-github/internal/server"
)

func nowTestFunc() time.Time {
	t, err := time.ParseInLocation("2006-01-02", "2023-12-12", time.Local)
	if err != nil {
		panic(err)
	}
	return t
}

func TestServe(t *testing.T) {
	var testCases = map[string]struct {
		input  string
		expect string
	}{
		"test Serve": {input: "", expect: "example serve command"},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			s := server.New(nowTestFunc)
			result := s.Serve()
			assert.Equal(tt, tc.expect, result, name)
		})
	}
}

func TestGreet(t *testing.T) {
	var testCases = map[string]struct {
		input       []string
		expect      string
		expectError bool
	}{
		"boomer": {
			input:  []string{"Jim", "1963-02-03"},
			expect: "Hello Jim, happy belated birthday for 312 days ago.",
		},
		"the doctor": {
			input:       []string{"Who", "2963-02-03"},
			expect:      "time travel detected",
			expectError: true,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			s := server.New(nowTestFunc)
			result, err := s.Greet(tc.input[0], tc.input[1])
			if tc.expectError {
				assert.EqualError(tt, err, tc.expect, name)
			} else {
				assert.NoError(tt, err, name)
				assert.Equal(tt, tc.expect, result, name)
			}
		})
	}
}

func FuzzGreet(f *testing.F) {
	f.Add("Joe", "2020-04-02")
	f.Fuzz(func(t *testing.T, name, birthDate string) {
		s := server.New(nowTestFunc)
		out, err := s.Greet(name, birthDate)
		if err != nil && out != "" {
			t.Errorf("%q, %v", out, err)
		}
	})
}
