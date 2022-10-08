package vector

import (
	"fmt"
	"math"
	"testing"
)

const DEG = 360 / (2.0 * math.Pi)
const RAD = 1.0 / DEG

func TestAngle(t *testing.T) {
	fmt.Println(V2D(1, 1).RightAngleToXAxis() * DEG)
	fmt.Println(V2D(1, -1).RightAngleToXAxis() * DEG)
	fmt.Println(V2D(-1, -1).RightAngleToXAxis() * DEG)
	fmt.Println(V2D(-1, 1).RightAngleToXAxis() * DEG)

	// fmt.Println(V2D(2.1, -7.8).RightAngleToXAxis() * DEG)
	// fmt.Println(V2D(25.1, -7.8).RightAngleToXAxis() * DEG)

	// fmt.Println(V2D(-2.1, -7.8).RightAngleToXAxis() * DEG)
	// fmt.Println(V2D(-25.1, -7.8).RightAngleToXAxis() * DEG)

	// fmt.Println(V2D(-2.1, 7.8).RightAngleToXAxis() * DEG)
	// fmt.Println(V2D(-25.1, 7.8).RightAngleToXAxis() * DEG)
}

func TestAngleTo(t *testing.T) {
	fmt.Println(V2D(1, 1).RightAngleTo(V2D(1, 0)) * DEG)
	fmt.Println(V2D(1, 1).RightAngleTo(V2D(1, -1)) * DEG)
	fmt.Println(V2D(1, 1).RightAngleTo(V2D(-1, -1)) * DEG)
	fmt.Println(V2D(1, 1).RightAngleTo(V2D(-1, 1)) * DEG)
}

func TestRotate(t *testing.T) {
	fmt.Println(V2D(1, 1).RotateBy(RAD * 45))
	fmt.Println(V2D(1, 1).RotateBy(-RAD * 45))
	//fmt.Println(V2D(1, 1).RightAngleTo(V2D(1, -1)) * DEG)
}
