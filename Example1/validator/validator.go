package validator

import "strings"

func IsValid(s string) bool {
	if len(s) < 6 { // Rule 1: at least 6 characters long
		return false
	}

	var repeatCount, seqCount int

	for i := 0; i < len(s)-1; i++ {

		switch {
		case s[i] == s[i+1]: // check for repeating characters
			repeatCount++
			if repeatCount >= 2 { // Rule 3
				return false
			}
		default:
			repeatCount = 0
		}

		switch {
		case i < len(s)-2 && s[i+2]-s[i+1] == s[i+1]-s[i] && s[i+2]-s[i+1] != 0: // check for sequence
			seqCount++
			if seqCount >= 2 { // Rule 2
				return false
			}
		default:
			seqCount = 0
		}
	}

	// Check for repeating two digit sequences
	if isValidTwoDigitSequence(s) { // Rule 4
		return false
	}
	return true
}

func isValidTwoDigitSequence(s string) bool {
	for i := 0; i < len(s)-3; i += 2 {
		seq := s[i : i+2]
		remaining := s[i+2:]
		if strings.Count(remaining, seq) > 1 {
			return false
		}
	}
	return true
}
