package main

import (
	"fmt"
	td "gogeom/TwoDimension"
	"math"
)

func main() {
	l1 := td.Line{Slope: 1, Intercept_x: 10.0, Intercept_y: 6.0}
	l2 := td.Line{Slope: 5, Intercept_x: 10.0, Intercept_y: math.Inf(1)}
	p1 := td.Point{X: 1, Y: 0}
	p2 := td.Point{X: 10, Y: 0}

	fmt.Println(l1, p1, p2)
	fmt.Println("Dist between p1 and p2", p1.Dist(&p2))

	fmt.Println("p1 lies on l1? : ", p1.LiesOnLine(&l1))
	fmt.Println("p2 lies on l1? : ", p2.LiesOnLine(&l1))
	fmt.Println("p1 lies on l2? : ", p1.LiesOnLine(&l2))
	fmt.Println("p2 lies on l2? : ", p2.LiesOnLine(&l2))

	p3 := td.Point{X: 10, Y: 0}

	fmt.Println("\np3 lies on l1? : ", p3.LiesOnLine(&l1)) // should return true, returns false :(
	fmt.Println("Dist of p2 from Origin : ", p2.FromOrigin())

	p4 := td.Point{X: 0, Y: 1}
	fmt.Println(td.MakeLine(&p1, &p4))

	fmt.Println(td.TriangleArea(&p1, &p4, &td.Origin))

}
