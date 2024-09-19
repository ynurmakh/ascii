package implementation

import (
	"Adlet/main_function/check"
	"Adlet/main_function/error_handling"
	"Adlet/main_function/implementation/support"
)

func Main_Operation(string_stream, banner string, permission bool) bool {
	if (check.Check_banner(banner) || check.Language_Checking(string_stream)) && error_handling.ErrorHandling("Main_Operatoin", nil, true, "") {
		return true
	}

	Obrabotka(banner)

	string_stream = support.Modification(string_stream)
	// fmt.Println("After Obrabotka", []byte(string_stream))
	if permission {
		TransferStringToGraph(string_stream)
	}
	return false
}
