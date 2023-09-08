// task #2 from https://golangify.com/concurency
package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
    for _, v := range []string{"hello world", "a bad paddle"} {
        downstream <- v
    }
    close(downstream)
}

func removeDublicates(upstream, downstream chan string) {
	for v := range upstream {
		for _, word := range strings.Fields(v) {
			downstream <- word
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go removeDublicates(c0, c1)
	printGopher(c1)
}

