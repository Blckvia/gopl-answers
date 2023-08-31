// ex2.4: compare popcount implementations, including looping table lookups and
// shift value.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x >>= 1
	}
	return count
}

func main() {
	for _, num := range os.Args[1:] {
		x, err := strconv.ParseUint(num, 10, 64)
		if err != nil {
			log.Printf("popcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(PopCount(x))
	}
}
