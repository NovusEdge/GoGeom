package gogeom

//Ellipse defines an ellipse on a 2D cartesian plane
type Ellipse struct {
	Focus_1      Point
	Focus_2      Point
	Major_Radius float64
	Minor_Radius float64
}

//MajorAxis returns the Line representing the major axis of the ellipse
func (e *Ellipse) MajorAxis() Line {
	return MakeLine(&e.Focus_1, &e.Focus_2)
}

//Center returns the Point representing the centre of the ellipse
func (e *Ellipse) Center() Point {
	return Section(&e.Focus_1, &e.Focus_2, 1)
}

//LatusRectum reports the length of the latus-rectum of the parabola
func (e *Ellipse) LatusRectum() float64 {
	return 2 * e.Minor_Raduis * e.Minor_Radius / e.Major_Raduis
}

//Eccentricity reports the eccentricity of the ellipse
func (e *Ellipse) Eccentricity() float64 {
	return e.Centre().Dist(e.Focus_1) / e.Major_Radius
}

//IsLine reports if the ellipse is a special case where it's just a line
func (e *Ellipse) IsLine() bool {
	return e.Focus_1.Dist(&e.Focus_2) == 2*e.Major_Radius
}

//IsCircle reports if the ellipse is a circle
func (e *Ellipse) IsCircle() bool {
	return e.Focus_1.IsEqual(e.Focus_2)
}
