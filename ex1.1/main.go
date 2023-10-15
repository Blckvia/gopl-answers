// ex1.1 prints command line arguments.
package main

import (
	"fmt"
	"os"
)

func a(a uint) {
  fmt.Println("isasaa")
}


func main() {
	fmt.Println(os.Args)
  a(5)
}

