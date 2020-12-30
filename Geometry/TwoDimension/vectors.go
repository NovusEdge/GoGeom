package gogeom

import (
	"math"
)

//Direction Defines the direction in which a vector is supposed to be in.
//type Direction struct {
//	Point_1 Point
//	Point_2 Point
//}

const I Vector = Vector{Point_1: Origin, Point_2: Point{X: 1, Y: 0}}
const J Vector = Vector{Point_1: Origin, Point_2: Point{X: 0, Y: 1}}

//Vector defines a 2D vector in direction from [Point_1] to [Point_2]
type Vector struct {
	Point_1 Point
	Point_2 Point
}

//CmpVector defines a 2D vector with x and y components as [X_Component] and [Y_Component]
type CmpVector struct {
	X_Component float64
	Y_Component float64
}

//Magnitude reports the magnitude of the vector
func (v *Vector) Magnitude() float64 {
	return v.Point_1.Dist(&v.Point_2)
}

func (cv *CmpVector) Magnitude() float64 {
	return math.Sqrt(cv.X_Component*cv.X_Component + cv.Y_Component*cv.Y_Component)
}

//Resize returns a vector resized with respect to [resize_factor]
func (v *Vector) Resize(resize_factor float64) Vector {
	return Vector{
		Point_1: Point{
			X: v.Point_1.X * resize_factor,
			Y: v.Point_1.Y * resize_factor,
		},
		Point_2: Point{
			X: v.Point_2.X * resize_factor,
			Y: v.Point_2.Y * resize_factor,
		},
	}
}

func (cv *CmpVector) Resize(resize_factor float64) CmpVector {
	return CmpVector{
		X_Component: cv.X_Component * resize_factor,
		Y_Component: cv.Y_Component * resize_factor,
	}
}

//UnitVector reports the unit vector along a vector
func (v *Vector) UnitVector() Vector {
	p1 := v.Point_1
	p2 := Point{
		X: v.Point_2.X / v.Magnitude,
		Y: v.Point_2.Y / v.Magnitude,
	}
	return Vector{Point_1: p1, Point_2: p2}
}

func (cv *CmpVector) UnitVector() CmpVector {
	return CmpVector{
		X_Component: cv.X_Component / cv.Magnitude(),
		Y_Component: cv.Y_Component / cv.Magnitude(),
	}
}

//DotProduct reports the dot product between 2 vectors
func (v1 *Vector) DotProduct(v2 *CmpVector) float64 {
	x_cmp1, y_cmp1 := v1.Point_2.X-v1.Point_1.X, v1.Point_2.Y-v1.Point_1.Y
	x_cmp2, y_cmp2 := v2.Point_2.X-v2.Point_1.X, v2.Point_2.Y-v2.Point_1.Y
	return x_cmp1*x_cmp2 + y_cmp1*y_cmp2
}

func (cv1 *CmpVector) DotProduct(cv2 *CmpVector) float64 {
	return cv1.X_Component*cv2.X_Component + cv1.Y_Component*cv2.Y_Component
}

//Angle reports the angle between 2 vector
func (v1 *Vector) Angle(v2 *Vector) float64 {
	return math.Acos(v1.DotProduct(v2) / (v1.Magnitude() * v2.Magnitude()))
}

func (cv1 *CmpVector) Angle(cv2 *CmpVector) float64 {
	return math.Acos(cv1.DotProduct(cv2) / (cv1.Magnitude() * cv2.Magnitude()))
}
