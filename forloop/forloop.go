package main

import "fmt"

func main() {

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("i: %v\n", i)
	// }

	//#########################################
	//forever and while loop
	// j := 0
	// for true { // j < m5
	// 	fmt.Printf("j: %v\n", j)
	// }
	//#########################################

	//#########################################
	// for len = range
	num := [3]string{"d", "a", "y"}

	// for i := 0; i < len(num); i++ {
	// 	fmt.Printf("ii: %v\n", i)
	// }

	for i := range num {
		fmt.Printf("i: %v\n", i)
	}

	fmt.Println("#################################")
	for i, v := range num {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
	//#########################################
}
