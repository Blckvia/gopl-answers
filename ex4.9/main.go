// ex4.9 counts word frequency for stdin.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	frequency := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		frequency[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "frequency: %v", err)
		os.Exit(1)
	}
	for word, n := range frequency {
		fmt.Printf("%-30s %d\n", word, n)
	}
}