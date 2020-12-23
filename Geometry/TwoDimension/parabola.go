package gogeom

//import (
//	"math"
//)

type Parabola struct {
	Focus     Point
	Directrix Line
}

//Axis returns the axis of the parabola
func (pb *Parabola) Axis() Line {

	F := pb.Focus
	D := pb.Diretrix

	in_y := D.Intercept_y - (F.X * (D.Slope*D.Slope + 1) / D.Slope)
	in_x := D.Slope * in_y

	return Line{
		Slope:       -1 / D.Slope,
		Intercept_x: in_x,
		Intercept_y: in_y,
	}
}

//Vertex returns the vertex of the parabola
func (pb *Parabola) Vertex() Point {

	D := pb.Diretrix
	A := pb.Axis()

	F := pb.Focus
	P := D.Intersection(&A)

	x := (F.X + P.X) / 2
	y := (F.Y + P.Y) / 2

	return Point{
		X: x,
		Y: y,
	}
}

//LiesOnParabola reports if a point lies on a parabola
func (p *Point) LiesOnParabola(pb *Parabola) bool {
	d_Dist := pb.Diretrix.FromPoint(p)
	f_Dist := p.Dist(&pb.Focus)
	return f_Dist == d_Dist
}
