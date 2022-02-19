package square

import (
	"testing"
)

func TestSquareTriangle(t *testing.T) {
	if CalcSquare(10.0, SidesTriangle) != 43.30127018922193 {
		t.Error("Expected 43.30127018922193, got ", CalcSquare(10.0, SidesTriangle))
	}
}

func TestSquareSquare(t *testing.T) {
	if CalcSquare(10.0, SidesSquare) != 100.0 {
		t.Error("Expected 100.0, got ", CalcSquare(4.0, SidesSquare))
	}
}

func TestSquareCircle(t *testing.T) {
	if CalcSquare(10.0, SidesCircle) != 314.1592653589793 {
		t.Error("Expected 314.1592653589793, got ", CalcSquare(10.0, SidesCircle))
	}
}
