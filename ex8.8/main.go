// ex8.8 Reverb2 is a TCP server that simulates an echo using sync.WaitGroup and select with timeout
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	wg := sync.WaitGroup{}
	input := bufio.NewScanner(c)
	timeout := time.NewTimer(10 * time.Second)
	text := make(chan string)

	go func() {
		for input.Scan() {
			text <- input.Text()
		}
		close(text)
	}()

	for {
		select {
		case t, ok := <-text:
			if ok {
				wg.Add(1)
				timeout.Reset(10 * time.Second)
				go func() {
					defer wg.Done()
					echo(c, t, 1*time.Second)
				}()
			} else {
				wg.Wait()
				c.Close()
				return
			}
		case <-timeout.C:
			timeout.Stop()
			c.Close()
			fmt.Println("time is expired")
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
