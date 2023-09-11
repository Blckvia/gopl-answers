package rotate

import (
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		s []int
		integer int
		want []int
	} {
		{[]int{1,2,3,4,5}, 2, []int{4,5,1,2,3}},
	}

	for _, test := range tests {
		ans := rotate(test.s, test.integer)
		for i, _ := range ans {
			if ans[i] != test.want[i] {
				t.Errorf("Slice mismatch:\nArr: %q, Integer: %q\nGot: %v\nWant: %v", test.s, test.integer, ans, test.want)
			}
		}
	}
}