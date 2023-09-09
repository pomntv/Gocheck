package main

import "fmt"

type course struct {
	name, instructor string
	price            float64
}

func discount(c course) float64 {
	// fmt.Printf("c.price: %v\n", c.price)
	return c.price - 100
}

func (c course) discount_M() float64 {
	return c.price - 100
}

func (c course) discount_M_realused(a float64) float64 {
	return c.price - a
}

func main() {
	a := course{name: "A kung", instructor: "", price: 500}

	fmt.Printf("discount aa: %#v\n", discount(a))

	fmt.Printf("a.discount_M(): %v\n", a.discount_M())

	fmt.Printf("a.discount_M_realused(200): %v\n", a.discount_M_realused(200))

}
