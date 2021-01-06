package epsslice

import "github.com/user/EPSUtils/code/EPSFloat64"

// Float64 is a slice of ints
type Float64 []float64

// Sum returns the sum of all elements in the slice
func (g Float64) Sum() float64 {
	var s float64
	for _, v := range g {
		s += v
	}
	return s
}

// Max returns the max of all elements in the slice
func (g Float64) Max() float64 {
	s := float64(0)
	for _, v := range g {
		s = EPSFloat64.Max(s, v)
	}
	return s
}
