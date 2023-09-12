// ex4.7 reverse a utf8 string
package reverse

import "unicode/utf8"

func reverse(arr []byte) {
	for i, j := 0, len(arr) - 1; i< j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reverseUtf (arr []byte) []byte {
	for i := 0; i < len(arr); {
		_, size := utf8.DecodeRune(arr[i:])
		// we have to reverse a letter. Without this reverse we will get wrong result
		// e.x: "hello" will be "olleh"
		reverse(arr[i: i+size])
		i += size
	}
	reverse(arr)
	return arr
}