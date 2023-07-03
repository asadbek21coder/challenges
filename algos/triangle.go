package algos

import "math"

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	p := (t.A + t.B + t.C) / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func (t Triangle) Perimetre() float64 {
	per := t.A + t.B + t.C
	return per
}
