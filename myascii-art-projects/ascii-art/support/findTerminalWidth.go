package support

import (
	"os/exec"
	"strconv"

	"Adlet/main_function/error_handling"
)

func Find_Terminal_Width() (bool, int) {
	widthCmd := exec.Command("tput", "cols")
	widthOut, err := widthCmd.Output()
	if error_handling.ErrorHandling("Find_Terminal_Width", err, false, "") {
		return true, 0
	}

	widthOut = widthOut[:len(widthOut)-1]
	width_terminal, err := strconv.Atoi(string(widthOut))
	if error_handling.ErrorHandling("Find_Terminal_Width", err, false, "") {
		return true, 0
	}
	return false, width_terminal
}
