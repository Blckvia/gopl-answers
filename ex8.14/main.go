// ex8.14 is a chat server that prompts clients for a name upon connection.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// !+broadcaster
type client struct {
	outgoing chan<- string // an outgoing message channel
	name     string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.outgoing <- msg
			}

		case cli := <-entering:
			clients[cli] = true

			for c := range clients {
				if c != cli {
					cli.outgoing <- cli.name
				}
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.outgoing)
		}
	}
}

//!-broadcaster

// !+handleConn
func handleConn(conn net.Conn) {
	var cli client
	var who string
	ch := make(chan string) // outgoing client messages
	out := make(chan string, 10)
	in := make(chan string)
	go clientWriter(conn, in)

	cli.outgoing = ch
	cli.name = who

	timer := time.NewTimer(5 * time.Minute)
	out <- "Enter your name:"
	select {
	case name := <-in:
		who = name
	case <-timer.C:
		conn.Close()
		return
	}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	timeout := time.NewTimer(5 * time.Minute)
	go func() {
		<-timeout.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timeout.Reset(5 * time.Minute)
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

// !+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
