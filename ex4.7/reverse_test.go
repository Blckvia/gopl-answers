package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct{
		s, want []byte
	}{
		{[]byte("kerangaR"), []byte("Ragnarek")},
	}

	for _, test := range tests {
		ans := reverseUtf(test.want)
		for i := range ans {
			if ans[i] != test.want[i] {
				t.Errorf("Arr: %q, Got: %v Want: %v", test.s, ans, test.want)
			}
		}
	}
}