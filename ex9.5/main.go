// ex9.5 tests of performance of ping-ponging goroutines.
package main

import (
	"fmt"
	"time"
)

func main() {
	in, out := make(chan int), make(chan int)
	var counter int

	t := time.NewTimer(1 * time.Minute)
	done := make(chan struct{})
	shutdown := make(chan struct{})

	go func() {
	loop:
		for {
			select {
			case <-shutdown:
				break loop
			case v := <-in:
				counter++
				out <- v
			}
			done <- struct{}{}
		}
	}()

	go func() {
	loop:
		for {
			select {
			case <-shutdown:
				break loop
			case v := <-out:
				counter++
				in <- v
			}
			done <- struct{}{}
		}
	}()
	in <- 1

	select {
	case <-in:
	case <-out:
	}
	<-done
	<-done
	t.Stop()
	fmt.Printf("Amount of counter: %v\n", counter)
}
