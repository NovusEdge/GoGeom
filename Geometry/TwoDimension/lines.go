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
func MakeLine(p_1 *Point, p_2 *Point) (line1 Line) {
	slope := float64((p_2.Y - p_1.Y) / (p_2.X - p_1.X))
	in_y := (p_1.Y - slope*p_1.X)
	in_x := (p_1.X - (p_1.Y / slope))

	return Line{
		Slope:       slope,
		Intercept_x: in_x,
		Intercept_y: in_y,
	}
}

//LiesOnLine reports if a point lies on a line
func (p *Point) LiesOnLine(line1 *Line) bool {
	in_x := line1.Intercept_x
	in_y := line1.Intercept_y

	// if (p.X == in_x && p.Y == 0) || (p.X == 0 && p.Y == in_y) {
	// 	return true
	// }

	if in_x == math.Inf(1) || in_x == math.Inf(-1) {
		return p.X == in_x
	} else if in_y == math.Inf(1) || in_y == math.Inf(-1) {
		return p.Y == in_y
	} else {
		return p.Y == (line1.Slope*p.X)+in_y
	}
}

//IsPerp reports if a line is perpendicular to another
func (line1 *Line) IsPerp(line2 *Line) bool {
	return line1.Slope*line2.Slope == -1
}

//IsParallel reports if a line is parallel to another
func (line1 *Line) IsParallel(line2 *Line) bool {
	return line1.Slope == line2.Slope
}

//Angle reports the angle (in radians) of the line with the axes
func (line1 *Line) Angle() (float64, float64) {
	if (line1.Intercept_x == math.Inf(-1) || line1.Intercept_x == math.Inf(1)) && (line1.Intercept_y == math.Inf(-1) || line1.Intercept_y == math.Inf(1)) {
		return math.NaN(), math.NaN()
	}

	if line1.Intercept_x == math.Inf(-1) || line1.Intercept_x == math.Inf(1) {
		return math.NaN(), math.Pi / 2
	} else if line1.Intercept_y == math.Inf(-1) || line1.Intercept_y == math.Inf(1) {
		return math.Pi / 2, math.NaN()
	} else {
		x_angle := math.Atan(line1.Slope)
		return x_angle, (math.Pi / 2) - x_angle
	}
}

//LineAngle reports the angle formed between 2 lines
func (line1 *Line) LineAngle(line2 *Line) float64 {
	if line1.Slope == line2.Slope {
		return math.NaN()
	}
	return math.Atan((line1.Slope - line2.Slope) / (1 - line1.Slope*line2.Slope))
}

//FromPoint reports the perpendicular distance of the line from a point
func (line1 *Line) FromPoint(p *Point) float64 {
	if line1.Intercept_x == math.Inf(-1) || line1.Intercept_x == math.Inf(1) {
		return line1.Intercept_y
	} else if line1.Intercept_y == math.Inf(-1) || line1.Intercept_y == math.Inf(1) {
		return line1.Intercept_x
	} else {
		m := line1.Slope
		return math.Abs((p.X*m - p.Y + line1.Intercept_y) / math.Sqrt(m*m+1))
	}

}

//FromOrigin reports the perpendicular distance of the line from the origin
func (line1 *Line) FromOrigin() float64 {
	return line1.FromPoint(&Origin)
}

//FromLine reports the perpendicular distance between 2 parallel lines
func (line1 *Line) FromLine(line2 *Line) float64 {
	return math.Abs((line1.Intercept_y - line2.Intercept_y) / math.Sqrt(line1.Slope*line1.Slope+1))
}

//IsTangent reports if a line is tangential to circle [c]
func (line1 *Line) IsTangent(c *Circle) bool {
	return math.Abs(line1.Intercept_y) == math.Abs(c.Radius*math.Sqrt(1+line1.Slope*line1.Slope))
}

//Intersection reports the Point where 2 lines intersect
func (line1 *Line) Intersection(line2 *Line) Point {
	x := (line1.Intercept_y - line2.Intercept_y) / (line1.Slope - line2.Slope)
	y := line1.Slope*x + line1.Intercept_y

	return Point{
		X: x,
		Y: y,
	}
}

//TODO: complete this function...

//IntersectsCircle reports if a line intersects a circle, and at what point(s)
func (line1 *Line) IntersectsCircle(c *Circle) (bool, *Point, *Point) {
	if l1.IsTangent(c) {
		return true, nil, nil
	}
	return false, nil, nil
}
