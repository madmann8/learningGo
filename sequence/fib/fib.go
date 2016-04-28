//Needs work still
package fib

import (
	"fmt"
)

func Fib (n int) int {
	p:= 0
	q:= 1
	for i := 0; i < n; i++ {
		p, q = q, p+q
	}
	return p
}

func add


