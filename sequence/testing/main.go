package main

import (
	"github.com/madmann8/learningGo/sequence/fib"
	"fmt"
)

func main() {
	t:= fib.Fib{n:10, m:5}
	fmt.Println(t.Add())
}