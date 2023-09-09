package main

import "fmt"

func adder() (func() int, func() int) {
	sum := 0
	return func() int {
			sum = sum + 1
			return sum
		}, func() int {
			return sum
		}
}

func main() {

	inc, curr := adder()

	y := curr()
	fmt.Println("1 curr 1 :", y)

	v := inc()
	fmt.Println("2 inc 1 ", v)

	y = curr()
	fmt.Println("3 curr 2 :", y)

	v = inc()
	fmt.Println("4 inc 2 ", v)

	y = curr()
	fmt.Println("5 curr 3 ", y)
}
