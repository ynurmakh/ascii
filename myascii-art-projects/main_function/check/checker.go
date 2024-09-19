package check

import (
	"Adlet/main_function/error_handling"
)

func Check_banner(banner string) bool {
	if Comparing_hash(banner) {
		error_handling.ErrorHandling("Check_banner", nil, true, "")
		return true
	}
	return false
}
