// ex11.5 tests strings.Split using a table-driven test.
package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want int
	}{
		{"a:b:c", ":", 3},
		{"k,v,c,b", ",", 4},
	}
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if got, want := len(words), 3; got != want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.s, test.sep, got, test.want)
		}
	}
}
