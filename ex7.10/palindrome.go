// ex7.10 create function which will check is palindrom the interface
package palindrome

import (
	"sort"
)

// IsPalindrome determines whether the given sort.Interface is a palindrome.
//
// The function takes a sort.Interface as a parameter, which represents a collection of elements that can be sorted.
// The function returns a boolean value indicating whether the sort.Interface is a palindrome.
func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for j > i {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
