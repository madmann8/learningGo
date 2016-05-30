// Package fib allows for usage of various Fibonacci Numbers

package fib

import (
	"math"
)


type Fib struct {
	FibNum float64
}

var Phi float64

//Initiates Fibonacci Sequence
func NewFib() *Fib {
	var f = &Fib{
		FibNum: 0,
	}
	Phi = (1 + math.Sqrt(5)) / 2
	return f
}
//Fib N finds nth number of Fibonacci sequence

func (f *Fib) FibN(N float64) *Fib {
	f.FibNum = (((math.Pow(Phi, N)) - math.Pow((-1*Phi), (-1*N))) / math.Sqrt(5))
	return f
}

//Adds the N Fibonacci Values to the previously calculated number(s)

func (f *Fib) Add(N float64) *Fib {
	var r float64
	r=(((math.Pow(Phi, N)) - math.Pow((-1*Phi), (-1*N))) / math.Sqrt(5))
	f.FibNum += float64(r)
	return f
}

//Multiplys the N Fibonacci Values by the previously calculated number(s)

func (f *Fib) Mul(N float64) *Fib {
	var r float64
	r=(((math.Pow(Phi, N)) - math.Pow((-1*Phi), (-1*N))) / math.Sqrt(5))
	f.FibNum *= float64(r)
	return f
}

//Divides the N Fibonacci Values by the previously calculated number(s)

func (f *Fib) Div(N float64) *Fib {
	var r float64
	r=(((math.Pow(Phi, N)) - math.Pow((-1*Phi), (-1*N))) / math.Sqrt(5))
	f.FibNum /= r
	return f

}

//Adds all Fibonacci values prior to N and including N

func (f *Fib) Sum(N float64) float64 {
	var r float64
	r = 0
	var t float64
	for i := 1; float64(i) < N+1; i++ {
		t = ((math.Pow(Phi, float64(i))) - math.Pow((-1 * Phi), (-1 * float64(i)))) / math.Sqrt(5)
		r = t + r
	}
	return r
}
//Required to end sequence
func (f *Fib) End() float64 {
	return f.FibNum
}
//Returns Golden Ratio
func (f *Fib) Phi() float64 {
	return Phi
}
