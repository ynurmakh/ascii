package implementation

import (
	"Adlet/ascii-art/support"
	"Adlet/justify/support/function"
	"Adlet/main_function/error_handling"
	MF "Adlet/main_function/implementation"
	MFsupport "Adlet/main_function/implementation/support"
)

func Solve(first, second, third string) bool {
	first = first[8:]
	if !(first == "center" || first == "right" || first == "left" || first == "justify") && error_handling.ErrorHandling("Solve", nil, true, "There is a problem with keys") {
		return true
	}
	flag, terminal_width := support.Find_Terminal_Width()
	if flag && error_handling.ErrorHandling("Solve", nil, true, "") {
		return true
	}
	if MF.Main_Operation(second, third+".txt", false) && error_handling.ErrorHandling("Solve", nil, true, "") {
		return true
	}
	second = MFsupport.Modification(second)

	if first == "center" {
		function.Center(terminal_width, second, third)
	} else if first == "right" {
		function.Right(terminal_width, second, third)
	} else if first == "left" {
		function.Left(terminal_width, second, third)
	} else if first == "justify" {
		function.Justify(terminal_width, second, third)
	}

	return false
}
