package check

import (
	"Adlet/main_function/error_handling"
)

func Language_Checking(Sentence string) bool {
	for i := 0; i < len(Sentence); i++ {
		if !(Sentence[i] >= 7 && Sentence[i] <= 126) {
			error_handling.ErrorHandling("Language_Checking", nil, true, "Found non-ASCII character")
			return true
		}
	}
	return false
}
