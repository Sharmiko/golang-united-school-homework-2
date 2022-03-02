package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type Sides int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const SidesTriangle Sides = 3
const SidesSquare Sides = 4
const SidesCircle Sides = 0

// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	var res float64
	if sidesNum == 0 {
		res = math.Pi * sideLen * sideLen
	} else if sidesNum == 3 {
		res = (math.Sqrt(3) / 4) * (sideLen * sideLen)
	} else if sidesNum == 4 {
		res = sideLen * sideLen
	}

	return res
}
