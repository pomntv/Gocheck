package main

import "fmt"

type course struct {
	name, instructor string
	price            int
}

func (c *course) discount() int {
	// (*c).price = (*c).price - 500 //same
	c.price = c.price - 500 //same
	// fmt.Println("discount:", c.price)
	return c.price
}

func main() {
	c := new(course)
	// Alternatively, you can initialize c like this:
	// c := &course{"Basic Go pointer", "aaa", 7000}

	c.price = 5000
	fmt.Printf("c.price: %v\n", c.price)

	fmt.Println("price:", c.discount()) // same
}
