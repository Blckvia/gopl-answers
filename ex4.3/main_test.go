package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		s, want [3]string
	} {
		{[3]string{"kiss", "love", "hate"}, [3]string{"hate", "love", "kiss"}},
		{[3]string{"you", "love", "I"}, [3]string{"I", "love", "you"}},
	}

	for _, test := range tests {
		reverse(&test.s)
		if !reflect.DeepEqual(test.s, test.want) {
			t.Errorf("ans: %q, want: %q", test.s, test.want)
		}
	}
}