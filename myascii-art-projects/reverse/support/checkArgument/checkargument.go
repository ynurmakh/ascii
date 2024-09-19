package checkArgument

import (
	"os"

	"Adlet/main_function/error_handling"
)

func Check(arguments []string) (bool, string, string, string) {
	if !(len(arguments) > 0 && len(arguments) <= 2) && error_handling.ErrorHandling("Check", nil, true, "There should be more/equal than one or less/equal to three arguments.") {
		return true, "", "", ""
	}
	if len(arguments) == 1 {
		first := arguments[0]

		if len(first) > 0 && len(first) <= 10 {
			return false, "", first, "standard"
		}

		if first[:10] == "--reverse=" {
			file_name := first[10:]
			extension := file_name[len(file_name)-4:]
			if extension != ".txt" && error_handling.ErrorHandling("Check", nil, true, "The extension is not correct, supposed to be \".txt\"") {
				return true, "", "", ""
			}
			_, err := os.Open(file_name)
			if error_handling.ErrorHandling("Check", err, false, "") {
				return true, "", "", ""
			}
			return false, first, "", ""
		} else if first[0] == '-' && first[1] == '-' && error_handling.ErrorHandling("Check", nil, true, "Usage: go run . [OPTION]\nEX: go run . --reverse=<fileName>\n") {
			return true, "", "", ""
		}
		return false, "", first, "standard"
	}
	if len(arguments) == 2 {
		if !(arguments[1] == "standard" || arguments[1] == "shadow" || arguments[1] == "thinkertoy") && error_handling.ErrorHandling("Check", nil, true, "The entered banner is not correct") {
			return true, "", "", ""
		}
	}

	return false, "", arguments[0], arguments[1]
}
