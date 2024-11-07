package v2

import (
	
	"math"
)

type Trougao struct{
	A float64
	B float64
	C float64
}

func Heron(t *Trougao) float64{
	s := (t.A + t.B + t.C) /2
	return math.Sqrt(s*(s-t.A)*(s-t.B)*(s-t.C))
}