package main

import (
	"fmt"
)

func main()  {
	fmt.Println(Fib(3783783733))
}


func Fib (n int) int {
	p:= 0
	q:= 1
	for i := 0; i < n; i++ {
		p, q = q, p+q
	}
	return p
}


