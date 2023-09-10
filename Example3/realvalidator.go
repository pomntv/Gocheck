package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

// Password function as you defined
func Password(pass string) (bool, string) {
	var (
		upp, low, num, sym bool
		tot                uint8
	)
	fmt.Printf("Lenght Input Password: %v\n", len(pass))
	buffer_size := 3
	// buffer_sizeint := strconv.Itoa(buffer_size)
	fmt.Printf("buffer_size: %v\n", buffer_size)
	buffer := make([]rune, 0, buffer_size) // Initialize buffer to hold last buffer_size characters
	Buffer_append := ""
	for _, char := range pass {
		//#### Buffer Zone ################################################################################
		// Update the buffer to include the latest character
		buffer = append(buffer, char)
		if len(buffer) > buffer_size {
			// Remove the earliest character from the buffer
			buffer = buffer[1:]
		}

		// Print the characters in the buffer
		if len(buffer) == buffer_size {
			// fmt.Println("Buffer:", string(buffer)) // each window buffer
			Buffer_append = Buffer_append + string(buffer)
		}

		// fmt.Printf("All Buffer_append: %v \n", Buffer_append)

		//#################################################################################################
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false, "Invalid character found"
		}
	}

	//### End for loop ####################################################################################

	//### Extend buffer ####################################################################################
	var result []string
	for i := 0; i < len(Buffer_append); i += buffer_size {
		// Ensure not to go out of bounds
		if i+buffer_size > len(Buffer_append) {
			break
		}
		// Append substring to result slice
		result = append(result, Buffer_append[i:i+buffer_size])
	}
	fmt.Printf("result: %#v\n", result)
	//### Only num pass ################################################################################
	var passed []string
	for _, item := range result {
		_, err := strconv.Atoi(item)
		if err == nil {
			passed = append(passed, item)
		}
	}
	if passed != nil {
		fmt.Println("The Number passed strings are:", passed)
		result = passed
	}
	//#### check duplicate ################################################################################
	freqMap := make(map[string]int)
	for _, item := range result {
		freqMap[item]++
	}

	for k, v := range freqMap {
		if v > 1 {
			fmt.Printf("%s appears %d times\n", k, v)
		}
	}
	//#### check increase  ################################################################################
	for _, item := range result {
		// Convert string to digits
		var digits []int
		for _, digitRune := range item {
			digit, err := strconv.Atoi(string(digitRune))
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			digits = append(digits, digit)
		}

		// Compute and print differential
		var diff []int
		for i := 0; i < len(digits)-1; i++ {
			absDiff := int(math.Abs(float64(digits[i+1] - digits[i])))
			diff = append(diff, absDiff)
		}
		// fmt.Printf("The differential of %s is %v\n", item, diff)

		if diff[0] == 1 && diff[1] == 1 {
			fmt.Printf("case Increase and Decrease %s is %v\n", item, diff)
		}
	}

	//### check Upper,Lower,Emoji ########################################################################
	if !upp {
		return false, "Must include at least one uppercase letter"
	}
	if !low {
		return false, "Must include at least one lowercase letter"
	}
	if !num {
		return false, "Must include at least one number"
	}
	if !sym {
		return false, "Must include at least one special character"
	}
	if tot < 8 {
		return false, "Must be at least 8 characters long"
	}

	return true, ""
}

func main() {
	// Test the Password function with various passwords
	testPasswords := []string{
		"12345559876541122334455555",
		"Password123!",
		// "password",
		// "Password",
		// "12345678",
		// "!@#$%^&*",
	}

	for _, pass := range testPasswords {
		fmt.Printf("Testing password: '%s'\n", pass)
		valid, reason := Password(pass)
		if valid {
			fmt.Printf("Password '%s' is valid.\n", pass)
		} else {
			fmt.Printf("Password '%s' is invalid: %s\n", pass, reason)
		}
		fmt.Println("----------------------")
	}
}
