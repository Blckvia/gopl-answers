// ex3.10 inserts commas into integer strings given as command-line arguments,
// without using recursion.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	prefix := n % 3

	if n <= 3 {
		return s
	}

	if prefix == 0 {
		prefix = 3
	}

	buf.WriteString(s[:prefix])

	for i := prefix; i < n; i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
	
}

func main() {
	for i := 1; i < len(os.Args); i++{
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}