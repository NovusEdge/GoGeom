package gogeom

import (
	"math"
)

//Circle defines a circle on a 2-d cartesian plane
type Circle struct {
	Radius float64
	Center Point
}

//Area reports the area of the circle
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Circumference reports the perimeter/circumference of the circle
func (c *Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

//LiesOnCircle reports if the point lies on a circle [c]
func (p *Point) LiesOnCircle(c *Circle) bool {
	return c.Radius*c.Radius == p.X*p.X+p.Y*p.Y
}

//Tangent returns a Line that is tangential to the circle at the point [p]
<<<<<<< HEAD
func (c *Circle) Tangent(p *Point) *Line {
	if !p.LiesOnCircle(c) {
		return nil
	}
	slope := -(p.X / p.Y)
	in_y := c.Radius * math.Sqrt(1+slope*slope)
	in_x := c.Radius * math.Sqrt(1+1/(slope*slope))
	return &Line{
		Slope:       slope,
		Intercept_x: in_x,
		Intercept_y: in_y,
=======
func (c *Circle) Tangent(p *Point) Line {
	if p.LiesOnCircle(c) {
		slope := -(p.X / p.Y)
		in_y := c.Radius * math.Sqrt(1+slope*slope)
		in_x := c.Radius * math.Sqrt(1+1/(slope*slope))
		return Line{
			Slope:       slope,
			Intercept_x: in_x,
			Intercept_y: in_y,
		}
	} else {
		panic("Cannot draw a tangent through a point that's not on the circle")
>>>>>>> ee5a21ed00c40d3809a6b2d896c7db31eee4cb23
	}
}

//SecantArea reports the areas of major and minor secants formed when line [l] cuts through the circle [c]
func (c *Circle) SecantArea(l *Line) (float64, float64, error) {
	return 0.0, 0.0, nil
}
