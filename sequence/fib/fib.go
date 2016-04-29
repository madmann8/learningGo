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
	return p+fib.M
}

func (fib *Fib) Mul() int {
	p:= 0
	q:= 1
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
	}
	f:=p
	return f*fib.M
}

func (fib *Fib) Div() int {
	p:= 0
	q:= 1
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
	}
	return p/fib.M
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

func (fib *Fib) Phi()int {
	p:= 0
	q:= 1
	for i := 0; i < fib.N -1; i++ {
		p, q = q, p+q
	}
	r:=p
	p= 0
	q= 1
	for i := 0; i < fib.N; i++ {
		p, q = q, p+q
	}
	g:=p/r
	return g
}






