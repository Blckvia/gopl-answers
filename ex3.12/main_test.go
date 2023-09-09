package main

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s, s2 string
		want bool
	} {
		{"anna", "anna", true},
		{"a", "ana", false},
		{"ana", "ana", true},
	}

	for _, test := range tests {
		ans := isAnagram(test.s, test.s2)
		if ans != test.want {
			t.Errorf("First string %q, seconds string %q, ans: %v, want: %v", test.s, test.s2, ans, test.want)
		}
	}
}