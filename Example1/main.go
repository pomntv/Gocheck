package main

import (
	"ExampleValidator1/validator"
	"fmt"
)

// 1) at least 6 charactors long
// only >6
// 2) increase & decrese number is not allowed (at lease 3 chars)
//     ex: 123456, 5734514, 162987

// 3) repeat number (at lease 3 chars)
//     ex: 111111, 5718880, 1927333

// 4) repeat 2 digits but sequence in set
//     ex: 112233, 3144556647, 17887766

func main() {
	input := "987654"
	if validator.IsValid(input) {
		fmt.Printf("%s is Pass\n", input)
	} else {
		fmt.Printf("%s is not Pass\n", input)
	}

}
