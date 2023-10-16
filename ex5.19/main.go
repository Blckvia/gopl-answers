// ex5.19 function which do not contain return statement
// returns non-zero value using panic & recover
package main

import "fmt"

func weirdDummy() (ret string, err error) {
	defer func() {
		p := recover()
		ret = "anything"
		err = fmt.Errorf("dummy error: %v", p)
	}()

	panic("dummy panic")
}
func main() {
	fmt.Println(weirdDummy())
}
