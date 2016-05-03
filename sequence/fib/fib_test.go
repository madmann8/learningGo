package fib

import (
	"testing"
	"math"
)

type testnumbers struct {
	values     []int
	FibvaluenN int
	FibvaluenM int
	Add        int
	Mult       int
	Div        float64
	Sum int
}
func RoundDown(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}

var tests = []testnumbers{
	{[]int{1, 2}, 1, 1, 2, 1, 1,1},
	{[]int{6, 10}, 8, 55, 63, 440, RoundDown(0.14545454545,4,),20},
	{[]int{36, 42}, 1493352, 267914296, 282844648, 4000054745112192, RoundDown(0.05572809,4),39088168},
}

func TestFibValues(t *testing.T) {
	for i := 0; i < 3; i++ {
		s := tests[i]
		v := s.FibvaluenN
		pair := s.values
		x:=pair[0]
		y:=pair[1]
		l := Fib{N:x,M:y}
		if v != l.Fib() {
			t.Error(
				"For", "Fib()",
				"expected", v,
				"got", l.Fib(),
			)
		}

	}
}

func TestFibAdd(t *testing.T) {
	for i := 0; i < 3; i++ {
		s := tests[i]
		v := s.Add
		pair := s.values
		x:=pair[0]
		y:=pair[1]
		l := Fib{N:x,M:y}
		if v != l.Add() {
			t.Error(
				"For", "Add()",
				"expected", v,
				"got", l.Add(),
			)
		}

	}

}
func TestFibMul(t *testing.T) {
	for i := 0; i < 3; i++ {
		s := tests[i]
		v := s.Mult
		pair := s.values
		x:=pair[0]
		y:=pair[1]
		l := Fib{N:x,M:y}
		if v != l.Mul() {
			t.Error(
				"For", "Mul()",
				"expected", v,
				"got", l.Mul(),
			)
		}

	}

}

func TestFibDiv(t *testing.T) {
	for i := 0; i < 3; i++ {
		s := tests[i]
		v := s.Div
		pair := s.values
		x := pair[0]
		y := pair[1]
		l := Fib{N:x, M:y}
		f := RoundDown(l.Div(), 4)
		if v != f {
			t.Error(
				"For", "Div()",
				"expected", v,
				"got", f,
			)
		}

	}

}

func TestFibPhi(t *testing.T) {
	for i := 0; i < 3; i++ {
		s := tests[i]
		pair := s.values
		x := pair[0]
		y := pair[1]
		l := Fib{N:x, M:y}
		f := RoundDown(l.Phi(), 0)
		var e bool
		if f==1 {
			e=true
		}
		if f== 2{
			e=true
		}
		if e!=true   {
			t.Error(
				"For", "Phi()",
				"expected", "1 or 2",
				"got", f,
			)
		}

	}

}

func TestFibSum(t *testing.T) {
	for i := 0; i < 3; i++ {
		s := tests[i]
		v := s.Sum
		pair := s.values
		x:=pair[0]
		y:=pair[1]
		l := Fib{N:x,M:y}
		if v != l.Sum() {
			t.Error(
				"For", "Sum()",
				"expected", v,
				"got", l.Fib(),
			)
		}

	}
}
