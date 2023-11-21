package concat_test

import (
	concat "Desktop/Projects/gopl-answers/ex1.3"
	"strings"
	"testing"
)

var args = []string{"hello", "my", "7", "friend"}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat.Concat(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, "")
	}
}
