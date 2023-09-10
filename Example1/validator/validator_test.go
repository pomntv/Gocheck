package validator

import (
	"fmt"
	"os"
	"testing"
)

func TestIsValid(t *testing.T) {
	totalNumbers := 0
	validNumbers := 0

	// Create and open the output text file for logging valid numbers
	f, err := os.Create("valid_numbers.txt")
	if err != nil {
		t.Fatalf("Could not create text file: %v", err)
	}
	defer f.Close()

	for i := 0; i <= 999999; i++ {
		s := fmt.Sprintf("%06d", i)
		totalNumbers++
		if IsValid(s) {
			validNumbers++

			// Write the valid number to the text file
			_, err := f.WriteString(fmt.Sprintf("%s\n", s))
			if err != nil {
				t.Fatalf("Could not write to text file: %v", err)
			}
		}
	}

	percentageValid := float64(validNumbers) / float64(totalNumbers) * 100
	t.Logf("Total numbers: %d", totalNumbers)
	t.Logf("Valid numbers: %d", validNumbers)
	t.Logf("Percentage of valid numbers: %.2f%%", percentageValid)
}
