package gogeom

import (
	"math"
	"reflect"
)

// Point defines a point on a 2-d cartesian plane
type Point struct {
	X float64
	Y float64
}

var Origin Point = Point{0, 0}

//LiesOn reports if a point lies on a line
func (p *Point) LiesOn(l *Line) bool {
	in_x := reflect.ValueOf(l.Intercept_x)
	in_y := reflect.ValueOf(l.Intercept_y)

	if reflect.TypeOf(l.Intercept_y) == nil {
		return p.X == in_x.Float()
	} else if reflect.TypeOf(l.Intercept_x) == nil {
		return p.X == in_y.Float()
	} else {
		return p.Y == (l.Slope*p.X)+in_y.Float()
	}
}

//Dist reports the distance of one point from another
func (p_1 *Point) Dist(p_2 *Point) float64 {
	return math.Sqrt(math.Pow(p_1.X-p_2.X, 2) + math.Pow(p_1.Y-p_2.Y, 2))
}

// FromOrigin reports the distance of the point from the origin ( O = (0, 0) )
func (p *Point) FromOrigin() float64 {
	return p.Dist(&Origin)
}
