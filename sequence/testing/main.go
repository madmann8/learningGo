package main

import (
	"github.com/madmann8/learningGo/sequence/fib"
	"fmt"
)

func main() {
	s := fib.Fib{10,6}
	fmt.Println(s.Sum())
}