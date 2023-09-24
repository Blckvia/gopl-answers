// ex5.2 counts the frequency of different tags in an html document on stdin.
package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"os"
)

func freqTag(r io.Reader) (map[string]int, error) {
	freq := make(map[string]int, 0)

	z := html.NewTokenizer(os.Stdin)
	var err error

	for {
		if z.Next() == html.ErrorToken {
			break
		}
		name, _ := z.TagName()
		if len(name) > 0 {
			freq[string(name)]++
		}
	}
	if err != io.EOF {
		return freq, err
	}
	return freq, nil
}

func main() {
	freq, err := freqTag(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for tag, count := range freq {
		fmt.Printf("%4d %s\n", count, tag)
	}
}
