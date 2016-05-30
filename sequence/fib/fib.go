// Package fib allows for usage of various Fibonacci Numbers

package main

import ()
import (
	"math"
)


type Fib struct {
	FibNum float64
	Phi    float64
}

func NewFib() *Fib {
	var f = &Fib{
		FibNum: 0,
		Phi:    1.6,
	}
	f.Phi = (1 + math.Sqrt(5)) / 2
	return f

}

func (f *Fib) FibN(N float64) *Fib {
	f.FibNum = ((math.Pow(f.Phi, N)) - math.Pow((-1*f.Phi), (-1*N))/math.Sqrt(5))
	return f
}

//Adds the N and M Fibonacci Values

func (f *Fib) Add(N float64) *Fib {
	r := 0
	t := 1
	for i := 0; float64(i) < N; i++ {
		r, t = t, r+t
	}
	f.FibNum += float64(r)
	return f
}

//Multiplies the N and M Fibonacci Values

func (f *Fib) Mul(M float64) *Fib {

	r := 0
	t := 1
	for i := 0; float64(i) < M; i++ {
		r, t = t, r+t
	}
	f.FibNum *= float64(r)
	return f
}

//Divides the N and M Fibonacci Values

func (f *Fib) Div(N float64) *Fib {
	var r float64
	var t float64
	r = 0
	t = 1
	for i := 0; float64(i) < float64(N); i++ {
		r, t = t, r+t
	}
	f.FibNum /= r
	return f

}

//Adds all Fibonacci values prior to N and N

func (f *Fib) Sum(N float64) *Fib {
	p := 0
	q := 1
	s := 0
	for i := 0; float64(i) < float64(N); i++ {
		p, q = q, p+q
		s += p
	}
	return f
}

//Calculates the Golden Ratio using N and M
