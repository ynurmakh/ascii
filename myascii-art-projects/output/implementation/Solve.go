package implementation

import (
	"Adlet/main_function/container"
	"Adlet/main_function/error_handling"
	"Adlet/main_function/implementation"
	"Adlet/output/support/filemanipulation"
	Writetofile "Adlet/output/support/filemanipulation"
)

func Solve(file_name, second, third string) bool { // first, second, third string)

	if len(second) == 0 && error_handling.ErrorHandling("Solve", nil, true, "It is required to insert a string to here, in order to be outputted in a particular file") {
		return true
	}
	if file_name[len(file_name)-4:] != ".txt" {
		error_handling.ErrorHandling("Solve", nil, true, "The file's extension should be \".txt\"")
		return true
	}
	file_name = file_name[9:]
	Writetofile.OutputToFile(third, "banner.txt")
	implementation.Main_Operation(second, third, true)
	if filemanipulation.OutputToFile(container.Result, file_name) && error_handling.ErrorHandling("Solve", nil, true, "") {
		return true
	}
	// }

	return false
}
