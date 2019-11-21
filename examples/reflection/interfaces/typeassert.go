package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer([]byte("Hello"))
	fmt.Printf("%T is a %t stringer\n", b, isStringer(b))

	i := 123
	fmt.Printf("%T is a %t stringer\n", i, isStringer(i))
}

func isStringer(v interface{}) bool {
	_, ok := v.(fmt.Stringer)
	return ok
}
