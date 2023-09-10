package validator

import "fmt"

// 1) at least 6 charactors long
// only >6
// 2) increase & decrese number is not allowed (at lease 3 chars)
//     ex: 123456, 5734514, 162987

// 3) repeat number (at lease 3 chars)
//     ex: 111111, 5718880, 1927333

// 4) repeat 2 digits but sequence in set
//     ex: 112233, 3144556647, 17887766
func Validate(password string) bool {
	if len(password) < 6 { // Rule 1: at least 6 characters long

		return false
	}

	var repeatCount, seqCount int
	for i := 0; i < len(password)-1; i++ {
		switch {
		case password[i] == password[i+1]: // check for repeating characters
			repeatCount++
			if repeatCount > 3 {
				fmt.Println("repeatCount > 3")
			}
		// case password[i+2]-password[i+1] == password[i+1]-password[i]:
		// 	seqCount++
		// 	if seqCount > 3 {
		// 		fmt.Println("seqCount > 3")
		// 	}
		default:
			repeatCount = 0
			seqCount = 0
			fmt.Printf("seqCount: %v\n", seqCount)
		}

	}

	return true
}
