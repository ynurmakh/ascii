package implementation

import (
	"fmt"

	"Adlet/ascii-art/implementation"
	"Adlet/main_function/container"
	"Adlet/main_function/error_handling"
)

func Solve(stream_string, banner string) bool {
	
	if implementation.OutputToTerminal(stream_string, banner) && error_handling.ErrorHandling("Solve", nil, true, "") {
		return true
	}

	fmt.Print(container.Result)
	return false
}
