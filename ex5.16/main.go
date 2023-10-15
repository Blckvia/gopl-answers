// ex5.16 variadic string-join function
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(variadicJoin("/", "hi", "there"))
	fmt.Println(variadicJoin(",", "Hello", "World"))
}

func variadicJoin(sep string, strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	b := bytes.Buffer{}
	for _, str := range strings {
		b.WriteString(str)
		b.WriteString(sep)
	}
	return b.String()
}
