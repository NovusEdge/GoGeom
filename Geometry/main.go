package main

import (
	"fmt"
	ttd "gogeom/ThreeDimension"
	td "gogeom/TwoDimension"
	"math"
)

func main() {
	l := td.Line{1, 10.0, math.Inf(1)}
	p1 := td.Point{1, 0}
	p2 := td.Point{1, 1}

	fmt.Println(l, p1, p2)
	fmt.Println(p1.Dist(&p2))

	fmt.Println(p1.LiesOn(&l))
	fmt.Println(p2.LiesOn(&l))

	p3 := td.Point{10, 0}

	fmt.Println(p3.LiesOn(&l))
	fmt.Println(p2.FromOrigin())

	p4 := td.Point{0, 1}
	fmt.Println(td.MakeLine(&p1, &p4))

	fmt.Println(td.TriangleArea(&p1, &p4, &td.Origin))

	p_1 := ttd.Point{0, 0, 0}
	fmt.Println(p_1)
}
