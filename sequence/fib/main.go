package main

import (
	"fmt"
)
//Needs work still
func main()  {
	fmt.Println(Fib(378378733))
}


func Fib (n int) int {
	p:= 0
	q:= 1
	for i := 0; i < n; i++ {
		p, q = q, p+q
	}
	return p
}


