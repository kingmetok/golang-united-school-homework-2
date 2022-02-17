package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sides int

const (
	SidesCircle   sides = 0
	SidesTriangle sides = 3
	SidesSquare   sides = 4
)

func CalcSquare(sideLen float64, sidesNum sides) float64 {
	switch sidesNum {
	case SidesTriangle:
		return math.Sqrt(sideLen*sideLen*3) / 4
	case SidesSquare, SidesCircle:
		return sideLen * sideLen
	default:
		panic("Incorrect sides num")
	}
}
