package gogeom

import (
	"math"
)

// Defines a line with slope and intercepts
type Line struct {
	Slope       float64
	Intercept_x float64
	Intercept_y float64
}

//MakeLine returns a line that passes through [p_1] and [p_2]
func MakeLine(p_1 *Point, p_2 *Point) (l Line) {
	slope := float64((p_2.Y - p_1.Y) / (p_2.X - p_1.X))
	in_y := (p_1.Y - slope*p_1.X)
	in_x := (p_1.X - (p_1.Y / slope))

	return Line{
		Slope:       slope,
		Intercept_x: in_x,
		Intercept_y: in_y,
	}
}

//LiesOn reports if a point lies on a line
func (p *Point) LiesOnLine(l *Line) bool {
	in_x := l.Intercept_x
	in_y := l.Intercept_y

	if in_x == math.Inf(1) || in_x == math.Inf(-1) {
		return p.X == in_x
	} else if in_y == math.Inf(1) || in_y == math.Inf(-1) {
		return p.Y == in_y
	} else {
		return p.Y == (l.Slope*p.X)+in_y
	}
}

//IsPerp reports if a line is perpendicular to another
func (l1 *Line) IsPerp(l2 *Line) bool {
	return l1.Slope*l2.Slope == -1
}

//IsParallel reports if a line is parallel to another
func (l1 *Line) IsParallel(l2 *Line) bool {
	return l1.Slope == l2.Slope
}

//Angle reports the angle (in radians) of the line with the axes
func (l *Line) Angle() (float64, float64) {
	if (l.Intercept_x == math.Inf(-1) || l.Intercept_x == math.Inf(1)) && (l.Intercept_y == math.Inf(-1) || l.Intercept_y == math.Inf(1)) {
		return math.NaN(), math.NaN()
	}

	if l.Intercept_x == math.Inf(-1) || l.Intercept_x == math.Inf(1) {
		return math.NaN(), math.Pi / 2
	} else if l.Intercept_y == math.Inf(-1) || l.Intercept_y == math.Inf(1) {
		return math.Pi / 2, math.NaN()
	} else {
		x_angle := math.Atan(l.Slope)
		return x_angle, (math.Pi / 2) - x_angle
	}
}

//LineAngle reports the angle formed between 2 lines
func (l1 *Line) LineAngle(l2 *Line) float64 {
	if l1.Slope == l2.Slope {
		return math.NaN()
	}
	return math.Atan((l1.Slope - l2.Slope) / (1 - l1.Slope*l2.Slope))
}

//FromPoint reports the perpendicular distance of the line from a point
func (l *Line) FromPoint(p *Point) float64 {
	if l.Intercept_x == math.Inf(-1) || l.Intercept_x == math.Inf(1) {
		return l.Intercept_y
	} else if l.Intercept_y == math.Inf(-1) || l.Intercept_y == math.Inf(1) {
		return l.Intercept_x
	} else {
		m := l.Slope
		return math.Abs((p.X*m - p.Y + l.Intercept_y) / math.Sqrt(m*m+1))
	}

}

//FromOrigin reports the perpendicular distance of the line from the origin
func (l *Line) FromOrigin() float64 {
	return l.FromPoint(&Origin)
}

//FromLine reports the perpendicular distance between 2 parallel lines
func (l1 *Line) FromLine(l2 *Line) float64 {
	return math.Abs((l1.Intercept_y - l2.Intercept_y) / math.Sqrt(l1.Slope*l1.Slope+1))
}

//IsTangent reports if a line is tangential to circle [c]
func (l *Line) IsTangent(c *Circle) bool {
	return math.Abs(l.Intercept_y) == math.Abs(c.Radius*math.Sqrt(1+l.Slope*l.Slope))
}

func (l1 *Line) Intersection(l2 *Line) Point {
	x := (l1.Intercept_y - l2.Intercept_y) / (l1.Slope - l2.Slope)
	y := l1.Slope*x + l1.Intercept_y

	return Point{
		X: x,
		Y: y,
	}
}
