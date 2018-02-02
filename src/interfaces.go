package main

import (
	"math"
	"fmt"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64  {
	return math.Pi*c.radius*c.radius
}

func (c circle) perim() float64  {
	return 2 * math.Pi * c.radius
}

func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}


func main() {
	r := rect{width:3, height:4}
	c := circle{radius:5}

	measure(r)
	measure(c)
}
