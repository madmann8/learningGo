//Needs work still
package fib

import()

type Fib struct{
	n int
	m int
}

func  (fib *Fib) Fib() {
	p:= 0
	q:= 1
	for i := 0; i < fib.n; i++ {
		p, q = q, p+q
	}
	return p
}

func (fib *Fib) Add() {
	p:= 0
	q:= 1
	for i := 0; i < fib.n; i++ {
		p, q = q, p+q
	}
	return p+fib.m
}

func (fib *Fib) Mul() {
	p:= 0
	q:= 1
	for i := 0; i < fib.n; i++ {
		p, q = q, p+q
	}
	return p*fib.m
}

func (fib *Fib) Div() {
	p:= 0
	q:= 1
	for i := 0; i < fib.n; i++ {
		p, q = q, p+q
	}
	return p/fib.m
}

func (fib *Fib) Sum() {
	p:= 0
	q:= 1
	s:=0
	for i := 0; i < fib.n; i++ {
		p, q = q, p+q
		s+=p
	}
	return s
}

func (fib *Fib) Phi() {
	p:= 0
	q:= 1
	for i := 0; i < fib.n-1; i++ {
		p, q = q, p+q
	}
	r:=p
	p= 0
	q= 1
	for i := 0; i < fib.n; i++ {
		p, q = q, p+q
	}
	g:=p/r
	return g
}






