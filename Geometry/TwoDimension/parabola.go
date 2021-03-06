package gogeom

<<<<<<< HEAD
// import (
// 	"math"
// )

=======
//Parabola defines a parabola on a 2D cartesian plane
>>>>>>> ee5a21ed00c40d3809a6b2d896c7db31eee4cb23
type Parabola struct {
	Focus    Point
	Diretrix Line
}

//Axis returns the axis of the parabola
func (pb *Parabola) Axis() Line {

	F := pb.Focus
	D := pb.Directrix

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

	D := pb.Directrix
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
	d_Dist := (pb.Directrix).FromPoint(p)
	f_Dist := p.Dist(&pb.Focus)
	return f_Dist == d_Dist
}

<<<<<<< HEAD
// func (l *Line) IntersectsParabola(pb *Parabola) (bool, *Point, *Point) {
// 	A := pb.Axis()
//
// 	if l.Slope == A.Slope && l.Intercept_x == A.Intercept_x && l.Intercept_y == A.Intercept_y {
// 		v := pb.Vertex()
// 		return true, &v, nil
// 	}
//
// }
=======
//IntersectsParabola reports the points where line intersects the parabola
// func (l *Line) IntersectsParabola(pb *Parabola) (bool, Point, Point) {
// 	A := pb.Axis()
// 	if l.Slope == A.Slope && l.Intercept_x == A.Intercept_x && l.Intercept_y == A.Intercept_y {
// 		return true, pb.Vertex(), Point{math.Inf(1), math.Inf(1)}
// 	}

// }

//LatusRectum reports the length of the latus-rectum of the parabola
func (pb *Parabola) LatusRectum() float64 {
	v := pb.Vertex()
	return 4 * pb.Focus.Dist(&v)
}
>>>>>>> ee5a21ed00c40d3809a6b2d896c7db31eee4cb23
