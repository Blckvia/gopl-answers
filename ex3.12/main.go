// ex3.12 the function checks whether two words are anagrams
package main

import (
	"fmt"
)

func isAnagram(s, s2 string) bool {
	if len(s) !=  len(s2) {
		return false
	}

	arr := make([]rune, len(s))
	for i, r := range s {
		arr[i] = r
	 }

	 for i, letter := range s2 {
		if letter != arr[i] {
			return false
		}
	 }
	 return true
}

func main() {
	fmt.Println(isAnagram("ana", "ann"))
}