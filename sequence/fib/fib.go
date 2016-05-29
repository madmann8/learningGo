// Package fib allows for usage of various Fibonacci Numbers

package fib

import ()
import (
	"math"
)

//Specifies the corresponding Fibonacci Number for the N value inputted
type Fib struct {
	FibNum float64
	Phi float64
}

func NewSeq(s string) *Fib {
	var s = &Fib{
		FibNum: 0,
		Phi: 0,
	}

	return s
}

func (Fib *Fib) Fib(N float64) *Fib  {
	Fib.FibNum=(Fib.Phi^(N)-(-1*Fib.Phi)^(-1*N))/math.Sqrt(5)
	return Fib
}

//Adds the N and M Fibonacci Values

func (Fib *Fib) Add (N float64) *Fib {
	r:= 0
	t:= 1
	for i := 0; i < N; i++ {
		r, t = t, r+t
	}
	Fib.FibNum+=r
	return Fib
}

//Multiplies the N and M Fibonacci Values

func  (Fib *Fib) Mul(M float64) *Fib {
	
	r:= 0
	t:= 1
	for i := 0; i < M; i++ {
		r, t = t, r+t
	}
	Fib.FibNum*=r
	return Fib
}

//Divides the N and M Fibonacci Values

func (Fib *Fib) Div(N float64) *Fib {
	var r float64
	var t float64
	r= 0
	t= 1
	for i := 0; float64(i) < float64(N); i++ {
		r, t = t, r+t
	}
	Fib.FibNum/=r
	return Fib

}

//Adds all Fibonacci values prior to N and N

func (Fib *Fib) Sum(N float64) *Fib {
	p:= 0
	q:= 1
	s:=0
	for i := 0; i < N; i++ {
		p, q = q, p+q
		s+=p
	}
	return Fib
}

//Calculates the Golden Ratio using N and M

func (Fib *Fib) Phi() *Fib {
	Fib.Phi=(1+math.Sqrt(5))/2
	return Fib
}
