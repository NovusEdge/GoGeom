package main

import (
	"fmt"
	td "gogeom/TwoDimension"
)

func main() {
	l := td.Line{1, 10.0, nil}
	p1 := td.Point{1, 0}
	p2 := td.Point{0, 1}

	fmt.Println(l, p1, p2)
	fmt.Println(p1.Dist(&p2))
}
