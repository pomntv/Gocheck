package main

import (
	"fmt"
)

func main() {
	// var skills [3]string
	skills := [3]string{"", "", ""}
	fmt.Printf("skills: %v\n", skills)
	fmt.Printf("skills: %#v\n", skills) // check knowledge

	var skills2 [30]string = [30]string{"s", "s", "s"} // use {} at array assigned
	fmt.Printf("skills2: %v\n", skills2)

	fmt.Printf("len skills2: %v\n", len(skills2)) // use len

}
