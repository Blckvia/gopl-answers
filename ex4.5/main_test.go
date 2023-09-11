package main

import (
	"fmt"
	"testing"
)

func TestDeleteAdjacentDub(t *testing.T) {
	tests := []struct {
		s, want []string
	} {
		{[]string{"a", "a", "b", "", ""}, []string{"a", "b", ""}},
	}

	for _, test := range tests {
		ans := deleteAdjacentDub(test.s)
		for i, _ := range ans {
			if ans[i] != test.want[i] {
				fmt.Errorf("Slice mismatch:\nArr: %q, Got: %v\nWant: %v", test.s, ans, test.want)
			}
		}
	}
}