package main

import (
	"fmt"
	"unicode"
)

// Password function as you defined
func Password(pass string) (bool, string) {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
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
		"Password123!",
		"password",
		"Password",
		"12345678",
		"!@#$%^&*",
	}

	for _, pass := range testPasswords {
		valid, reason := Password(pass)
		if valid {
			fmt.Printf("Password '%s' is valid.\n", pass)
		} else {
			fmt.Printf("Password '%s' is invalid: %s\n", pass, reason)
		}
	}
}
