package main

import (
	"github.com/madmann8/learningGo/sequence/fib"
	"fmt"
)

func main() {
	s:= fib.Fib{1,2}
	fmt.Println(s.Add())
}