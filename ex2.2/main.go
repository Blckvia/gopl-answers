// ex2.2 prints measurements given on the command line or stdin in various
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Feet float64
type Meter float64

func (f Feet) String() string { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%g mt", m) }

func FToM(f Feet) Meter { return Meter(f * 0.3048) }
func MToF(m Meter) Feet { return Feet(m * 3.28083) }

func Convertion(t float64) {
	lengthFeet := Feet(t)
	lengthMeter := Meter(t) 
	fmt.Printf("%s = %s, %s = %s\n\n",
	lengthFeet, FToM(lengthFeet), lengthMeter, MToF(lengthMeter))
}

func main() {
	if len(os.Args[1:]) > 0 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "uc: %v\n", err)
				os.Exit(1)
			}
			Convertion(t)
		}
	} else {
		for {
			input := bufio.NewReader(os.Stdin)
			s, err := input.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "Convertion: %v\n", err)
				os.Exit(1)
			}
			t, err := strconv.ParseFloat(s[:len(s)-1], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Convertion: %v\n", err)
				os.Exit(1)
			}
			Convertion(t)
		}
	}
}