package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	tests := []struct {
		s, want string
	}{
		{"1", "1"},
		{"123", "123"},
		{"5241", "5,241"},
		{"3891273891273", "3,891,273,891,273"},
	}

	for _, test := range tests {
		ans := comma(test.s)
		if ans != test.want {
			t.Errorf("comma: %q, ans: %q, want: %q", test.s, ans, test.want)
		}
	}
}
