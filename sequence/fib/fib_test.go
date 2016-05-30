package fib

import (
	"testing"
	"math"
)

type testnumbers struct {
	ValueN float64
	FibValueN float64
	Add        float64
	Mult       float64
	Div        float64
	Sum float64
}

func Round64(x float64) float64 {
	i, f := math.Modf(x)
	if f >= 0.5 {
		return i + 1.0
	}
	return i
}

func RoundDown(input float64, places float64) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Floor(digit)
	newVal = round / pow
	return
}

var tests = []testnumbers{
	{1, 1, 2, 1, 1,1},
	{6, 8, 16, 64, 1,20},
	{8, 21, 42 , 441, 1,143},
}

func  TestFibValue(t *testing.T) {
	for i := 0; i < 3; i++ {
		s := tests[i]
		v := s.ValueN
		c:=s.FibValueN
		l:=NewFib()
		x:=Round64(l.FibN(v).End())
		if c != x {
			t.Error(
				"For", "FibN()",
				"expected", c,
				"got", x,2,
			)
		}

	}
}

func TestFibAdd(t *testing.T) {
	for i := 0; i < 3; i++ {
		s := tests[i]
		v := s.ValueN
		c:=s.Add
		l:=NewFib()
		if c != Round64(l.FibN(v).Add(v).End()) {
			t.Error(
				"For", "Add()",v,
				"expected", s.Add,
				"got", l.Add(v),
			)
		}

	}
}
func TestFibMul(t *testing.T) {
	for i := 0; i < 3; i++ {
		s := tests[i]
		v := s.ValueN
		c:=s.Mult
		l:=NewFib()
		if c != l.FibN(v).Mul(v).End() {
			t.Error(
				"For", "Add()",
				"expected", s.Mult,
				"got", l.Mul(v),
			)
		}

	}
}

func TestFibDiv(t *testing.T) {
	for i := 0; i < 3; i++ {
		for i := 0; i < 3; i++ {
			s := tests[i]
			v := s.ValueN
			c := s.Div
			l := NewFib()
			if c != l.FibN(v).Div(v).End() {
				t.Error(
					"For", "Div()",
					"expected", s.Div,
					"got", l.Div(v),
				)
			}

		}
	}
}


func TestFibSum(t *testing.T) {
	for i := 0; i < 3; i++ {
		s := tests[i]
		v := s.Sum
		c:=s.Sum
		l:=NewFib()
		if c != s.Sum {
			t.Error(
				"For", "Sum()",
				"expected", s.Sum,
				"got", l.Sum(v),
			)
		}

	}
}

