package main

import (
	"github.com/madmann8/learningGo/sequence/fib"
	"fmt"
)

func main() {
	s := fib.Fib{999,555}
	fmt.Println(s.Phi())
}