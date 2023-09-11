// ex4.6 replaceUnicodeSpaces squashes each rune of adjacent Unicode spaces in
// UTF-8 encoded []byte slice into a single ASCII space
package main

import (
	"unicode"
	"unicode/utf8"
)

func replaceUnicodeSpaces(input []byte) []byte {
    output := make([]byte, 0, len(input))
    inSpace := false

    for _, b := range input {
        r, size := utf8.DecodeRune([]byte{b})

        if unicode.IsSpace(r) {
            if !inSpace {
                output = append(output, ' ')
                inSpace = true
            }
        } else {
            output = append(output, input[:size]...)
            inSpace = false
        }
    }

    return output
}

func main() {
    input := []byte("Hello\u2000\u2001\u2002\u2003\u2004\u2005\u2006World")
    result := replaceUnicodeSpaces(input)
    println(string(result)) 
}