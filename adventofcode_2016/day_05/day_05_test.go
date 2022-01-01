package day_05

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	tests := []struct {
		doorID string
		pwd    string
	}{
		{
			doorID: "abc",
			pwd:    "18f47a30",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res := Password(test.doorID)
			if res != test.pwd {
				t.Fatalf("password: expect %q, got %q", test.pwd, res)
			}
		})
	}
}

func TestPassword2(t *testing.T) {
	tests := []struct {
		doorID string
		pwd    string
	}{
		{
			doorID: "abc",
			pwd:    "05ace8e3",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%02d", i), func(t *testing.T) {
			res := Password2(test.doorID)
			if res != test.pwd {
				t.Fatalf("password2: expect %q, got %q", test.pwd, res)
			}
		})
	}
}
