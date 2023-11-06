// ex8.10 is a web-mirroring tool use http.NewRequest
package main

import (
	"fmt"
	"log"
	"os"
)

// !+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string, cancel <-chan struct{}) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := Extract(url, cancel)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

// !+
func main() {
	worklist := make(chan []string)
	cancel := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
	}()

	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link, cancel)
				}(link)
			}
		}
	}
}

//!-
