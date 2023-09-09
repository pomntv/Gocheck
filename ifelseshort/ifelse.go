package main

import (
	"fmt"
	"math"
)

func main() {
	lim := 225.0
	v := math.Pow(10, 2)
	if v < lim {
		fmt.Println("× power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g. \n", v, lim)
	}

	//short
	if a := math.Pow(10, 2); a < lim {
		fmt.Println("× power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g. \n", v, lim)
	}
	// not know a

}
