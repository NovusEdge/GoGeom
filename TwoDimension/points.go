package gogeom

import (
	"math"
	//	"reflect"
)

// Point defines a point on a 2-d cartesian plane
type Point struct {
	X float64
	Y float64
}

//LiesOn reports if a point lies on a line
func (p *Point) LiesOn(l *Line) bool {
	in_x := reflect.ValueOf(l.intercept_x).Type()
	in_y := reflect.ValueOf(l.intercept_y).Type()

	if in_y.Type() == nil {
		return p.x == in_x
	} else if in_x.Type() == nil {
		return p.y == in_y
	} else {
		return p.y == (l.slope*p.x)+in_y
	}
}

//Dist reports the distance of one point from another
func (p_1 *Point) Dist(p_2 *Point) float64 {
	return math.Sqrt(math.Pow(p_1.X-p_2.X, 2) + math.Pow(p_1.Y-p_2.Y, 2))
}

//test
