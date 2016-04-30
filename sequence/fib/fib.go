//Needs work still
package fib

import()

type Fib struct{
	N int
	M int
}

func  (fib *Fib) Fib() int {
	p:= 0
	q:= 1
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
	}
	return p
}

func (fib *Fib) Add() int {
	p:= 0
	q:= 1
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
	}
	f:= 0
	s:= 1
	for i := 0; i < fib.N; i++ {
		f, s = f, f+s
	}
	return p+f
}

func (fib *Fib) Mul() int {
	p:= 0
	q:= 1
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
	}
	f:= 0
	s:= 1
	for i := 0; i < fib.N; i++ {
		f, s = f, f+s
	}
	return f*p
}

func (fib *Fib) Div() int {
	p:= 0
	q:= 1
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
	}
	f:= 0
	s:= 1
	for i := 0; i < fib.N; i++ {
		f, s = f, f+s
	}
	return p/f
}

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

func (fib *Fib) Phi()float64 {
	var r float64
	var p float64
	var q float64
	p= 0
	q= 1
	for i := 0; i < fib.N -1; i++ {
		p, q = q, p+q
	}
	r=p
	p= 0
	q= 1
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
	}
	var g float64
	g=1 +(r/p)
	return float64(g)
}






