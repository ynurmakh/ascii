package filemanipulation

import (
	"os"

	"Adlet/main_function/error_handling"
)

func OutputToFile(stream_string, file_name string) bool {
	// Create the file
	file, err := os.Create(file_name)
	if error_handling.ErrorHandling("OutputToFile", err, false, "") {
		return true
	}
	defer file.Close() // Ensure file is closed even if errors occur

	// Write the string to the file

	_, err = file.WriteString(stream_string)

	if error_handling.ErrorHandling("OutputToFile", err, false, "") {
		return true
	}

	return false
}
