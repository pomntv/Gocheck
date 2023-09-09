package main

import "fmt"

func main() {
	today := "sunday"
	switch today {
	case "sunday":
		fmt.Printf("1 today is %v\n", today)
		fallthrough //test
	case "thueday":
		fmt.Printf("2 today is %v\n", today)
	case "wednesday":
		fmt.Printf("3 today is %v\n", today)
	case "thusday", "friday":
		fmt.Printf("4 today is %v\n", today)
	}

	//no expression
	today = "sss"
	switch {
	case today == "ss":
	case today == "sss" || today == "ss":
		fmt.Printf("5 today is %v\n", today)
	}

	//short
	switch theday := "monday"; theday {
	case "sa":
		fmt.Printf("1 today is %v\n", today)
	case "s":
		fmt.Printf("1 today is %v\n", today)
	}
	// theday end

}
