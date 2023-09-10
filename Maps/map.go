package main

import (
	"fmt"
)

func main() {
	//slice
	// 	ss := []string{}
	// 	fmt.Printf("ss: %v\n", ss)
	// map

	// mm := map[string]int{} // same at use make
	mm := make(map[string]int)

	fmt.Printf("mm: %v\n", mm)
	mm["answer"] = 12
	fmt.Printf("mm: %#v\n", mm)

	mm["answer"] = 13

	fmt.Printf("answer v: %#v\n", mm["answer"])
	fmt.Printf("all v: %#v\n", mm)

	delete(mm, "answer")
	fmt.Printf("delete v: %#v\n", mm)

	// mm["answer"] = 14
	n, ok := mm["answer"]
	if ok {
		fmt.Printf("ok n: %v\n", n)
	} else {
		fmt.Printf("no n: %v\n", n)
	}
}
