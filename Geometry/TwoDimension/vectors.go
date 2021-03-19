package gogeom

import "math"

//VectorFunctions ...
type VectorFunctions interface {
	Magnitude() float64
	Add()
	Angle()
}

//PositionVector ...
/*
	X is the coefficient of the x-component of the vector
	Y is the coefficient of the y-component of the vector
*/
type PositionVector struct {
	X float64
	Y float64
}

//Vector ...
/*
	Tail is the start-point of the vector
	Head is the end-point of the vector
*/
type Vector struct {
	Tail PositionVector
	Head PositionVector
}

/*
O represents a 0 vector.

I represents the unit vector along x-axis

J represents the unti vector along y-axis
*/
const (
	O PositionVector = Vector{0, 0}
	I PositionVector = Vector{1, 0}
	J PositionVector = Vector{0, 1}
)

//TODO: complete this...

//Add returns the addition of 2 vectors
func (pv *PositionVector) Add(opv *PositionVector) (res PositionVector) {
	return
}

//Add returns the addition of 2 vectors
func (v *Vector) Add(ov *Vector) (res Vector) {
	return
}

//Magnitude reports the magnitude of the vector
func (pv *PositionVector) Magnitude() float64 {
	return math.Sqrt(pv.X*pv.X + pv.Y*pv.Y)
}

//Magnitude reports the magnitude of the vector
func (v *Vector) Magnitude() float64 {
	return
}
