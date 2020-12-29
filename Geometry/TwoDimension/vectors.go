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

func (v *Vector) DotProduct() Vector {

}
