//Needs work still
package fib

import(
)

type Fib struct{
	N int
	M int
}

//Specifies the corresponding Fibonacci Number for the N value inputted

func  (fib *Fib) Fib() int {
	var p int
	var q int
	p= 0
	q= 1
	for i := 0; int(i) <int(fib.N); i++ {
		p, q = q, p+q
	}
	return p
}

//Adds the N and M Fibonacci Values

func (fib *Fib) Add() int {
	p:= 0
	q:= 1
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
	}
	f:= p
	r:= 0
	t:= 1
	for i := 0; i < fib.M; i++ {
		r, t = t, r+t
	}
	return f+r
}

//Multiplies the N and M Fibonacci Values

func (fib *Fib) Mul() int {
	p:= 0
	q:= 1
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
	}
	f:= p
	r:= 0
	t:= 1
	for i := 0; i < fib.M; i++ {
		r, t = t, r+t
	}
	return f*r
}

//Divides the N and M Fibonacci Values

func (fib *Fib) Div() float64 {
	var p float64
	p= 0
	var q float64
	q= 1
	for i := 0; float64(i) < float64(fib.N); i++ {
		p, q = q, p+q
	}
	var f float64
	f= p
	var r float64
	var t float64
	r= 0
	t= 1
	for i := 0; float64(i) < float64(fib.M); i++ {
		r, t = t, r+t
	}
	return f/r

}

//Adds all Fibonacci values prior to N and N

func (fib *Fib) Sum()int {
	p:= 0
	q:= 1
	s:=0
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
		s+=p
	}
	return s
}

//Calculates the Golden Ratio using N and M

func (fib *Fib) Phi()float64 {
	var p float64
	p= 0
	var q float64
	q= 1
	for i := 0; float64(i) < float64(fib.N); i++ {
		p, q = q, p+q
	}
	var f float64
	f= p
	var r float64
	var t float64
	r= 0
	t= 1
	for i := 0; float64(i) < float64(fib.M); i++ {
		r, t = t, r+t
	}
	var g float64
	g=f+r
	var b float64
	b= g/r
	return float64(b)
}






