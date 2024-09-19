package support

import (
	"io/ioutil"

	"Adlet/main_function/error_handling"
)

func Reading_file(file_name string) (bool, string) {
	var Mainstring string
	data, err := ioutil.ReadFile(file_name)
	if error_handling.ErrorHandling("Reading_file", err, false, "") {
		return true, ""
	}
	Mainstring = string(data)
	return false, Mainstring
}
