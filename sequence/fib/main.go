//Needs work still
package main

import (
	"fmt"
)
func main()  {
	fmt.Println(Fib(378378733))
}

//Just basic program, still need to add methods etc.
func Fib (n int) int {
	p:= 0
	q:= 1
	for i := 0; i < n; i++ {
		p, q = q, p+q
	}
	return p
}


