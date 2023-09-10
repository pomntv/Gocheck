package main

import (
	"fmt"
	"unicode"
)

func main() {
	sampleRunes := []rune{'a', '1', '!', '.', '@', ' ', '-', '+'}

	for _, r := range sampleRunes {
		if unicode.IsPunct(r) {
			fmt.Printf("The rune '%c' is a punctuation character.\n", r)
		} else {
			fmt.Printf("The rune '%c' is NOT a punctuation character.\n", r)
		}
	}
}
