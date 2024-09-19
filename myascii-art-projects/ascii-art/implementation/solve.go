package implementation

import (
	"Adlet/ascii-art/support"
	"Adlet/main_function/error_handling"
	"Adlet/main_function/implementation"
	MFsupport "Adlet/main_function/implementation/support"
)

func OutputToTerminal(stream_string, banner string) bool {
	if implementation.Main_Operation(stream_string, banner, false) {
		error_handling.ErrorHandling("OutputToTerminal", nil, true, "")
		return true
	}

	flag, terminal_width := support.Find_Terminal_Width()
	// fmt.Println("The terminal size is:", terminal_width)
	if flag && error_handling.ErrorHandling("OutputToTerminal", nil, true, "") {
		return true
	}

	stream_string = MFsupport.Modification(stream_string)
	Output(terminal_width, stream_string)

	return false
}
