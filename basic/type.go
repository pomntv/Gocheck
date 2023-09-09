package main

import (
	"fmt"
)

var outside string = "outside"

func main() {
	{
		msg, age, price, cbool := "Go", 26, 100, true
		fmt.Printf("type: %T -- msg: %#v\n", msg, msg)       // %s if need "" use %q
		fmt.Printf("type: %T -- age: %#v\n", age, age)       // %d
		fmt.Printf("type: %T -- price: %#v\n", price, price) // %.2f
		fmt.Printf("type: %T -- cbool: %#v\n", cbool, cbool) // %t
	}
	// fmt.Printf("a: %v\n", a) // undefined

	// ascii use rune (int32) > char 8
	var r rune = 'a'
	fmt.Printf("type: %T -- r: %v\n", r, r)

}
