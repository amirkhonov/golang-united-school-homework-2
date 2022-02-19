package square

import (
	"math"
)

type sideInt int

const (
	SidesTriangle sideInt = 3
	SidesSquare   sideInt = 4
	SidesCircle   sideInt = 0
)

func CalcSquare(sideLen float64, sidesNum sideInt) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * sideLen * sideLen
	case 3:
		return sideLen * sideLen * math.Sqrt(3.0) / 4.0
	case 4:
		return sideLen * sideLen
	default:
		return 0.0
	}
}
