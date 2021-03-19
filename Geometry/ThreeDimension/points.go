package gogeom

import (
	"math"
)

//Point defines a Point in 3D space
type Point struct {
	X float64
	Y float64
	Z float64
}

var Origin Point = Point{0, 0, 0}

//Dist reports the distance of one point from another
func (p1 *Point) Dist(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2) + math.Pow(p1.Z-p2.Z, 2))
}

//IsEqual reports if 2 points are coincident points
func (p1 *Point) IsEqual(p2 *Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y && p1.Z == p2.Z
}

// FromOrigin reports the distance of the point from the origin ( O = (0, 0) )
func (p *Point) FromOrigin() float64 {
	return p.Dist(&Origin)
}

//AreCollinear reports if 3 points are collinear
func AreCollinear(p1 *Point, p2 *Point, p3 *Point) bool {
	return p1.X*(p2.Y*p3.Z-p2.Z*p3.Y)+
		p1.Z*(p3.Z*p2.Z-p3.X*p2.Y) ==
		p1.Y*(p3.Z*p2.X-p2.Z*p3.X)
}

func TriangleArea(p1 *Point, p2 *Point, p3 *Point) float64 {
	if AreCollinear(p1, p2, p3) {
		return 0
	}

	a := p1.Dist(p2)
	b := p2.Dist(p3)
	c := p1.Dist(p3)

	s := (a + b + c) / 2

	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

//Section reports the point which divides the line joining [p1] and [p2] in the ratio [ratio]
func Section(p1 *Point, p2 *Point, ratio float64) Point {
	if ratio >= 1 {
		return Point{
			X: (ratio*p1.X - p2.X) / (ratio - 1),
			Y: (ratio*p1.Y - p2.Y) / (ratio - 1),
			Z: (ratio*p1.Z - p2.Z) / (ratio - 1),
		}
	}

	return Point{
		X: (ratio*p1.X + p2.X) / (ratio + 1),
		Y: (ratio*p1.Y + p2.Y) / (ratio + 1),
		Z: (ratio*p1.Z + p2.Z) / (ratio + 1),
	}

}
