package check

import (
	"bufio"
	"crypto/sha256"
	"io"
	"os"

	"Adlet/main_function/error_handling"
)

func CalculateSHA256Hash(filePath string) ([]byte, bool) {
	file, err := os.Open(filePath)
	if error_handling.ErrorHandling("CalculateSHA256Hash", err, false, "") {
		return []byte{}, true
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if error_handling.ErrorHandling("CalculateSHA256Hash", err, false, "") {
		return []byte{}, true
	}

	return hash.Sum(nil), false
}

func Comparing_hash(filePath string) bool {
	current_file_name := "../main_function/banner/" + filePath
	first, flag := CalculateSHA256Hash(current_file_name)
	if flag {
		error_handling.ErrorHandling("Comparing_has", nil, true, "")
		return true
	}

	original_file_name := "../main_function/check/file_key/" + filePath
	file, err := os.Open(original_file_name)
	if error_handling.ErrorHandling("Comparing_has", err, false, "") {
		return true
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}
	second := []byte(text)

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			filePath = "There is mismatch between original file and current file. The banner is: " + filePath
			error_handling.ErrorHandling("Comparing_hash", nil, true, filePath)
			return true
		}
	}

	return false
}
