package main

import (
	"github.com/madmann8/learningGo/sequence/fib"
	"fmt"
)

func main() {
	s:= fib.ini{36,2}
	fmt.Println(s.Sum())
}