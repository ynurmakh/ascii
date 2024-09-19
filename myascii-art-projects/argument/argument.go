package argument

import (
	"Adlet/main_function/error_handling"
)

/*
The must  have not been banner format is:

	string string string
	banner (string, banner)
	option string string
	option banner string
	option (in some cases)
	string string
	banner string
*/
func option_retrieve(option string) string {
	var object string = ""
	for i := len(option) - 1; i >= 0; i-- {
		if option[i] != '=' {
			object = string(option[i]) + object
		} else {
			break
		}
	}
	return object
}

func option_identification(option string) (bool, string) {
	if len(option) < 2 {
		return false, ""
	}
	if !(option[0] == '-' && option[1] == '-') {
		return false, ""
	}

	file_name := option_retrieve(option)
	option = option[:len(option)-len(file_name)]

	var object string = ""
	for i := 2; i < len(option); i++ {
		if option[i] != '=' {
			object += string(option[i])
		} else {
			break
		}
	}
	if !(object == "reverse" || object == "align" || object == "output" || object == "color") {
		return false, ""
	}

	return true, object
}

func banner_identification(banner string) bool {
	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" {
		return true
	}
	return false
}

func Check_letter(stream_string string) bool {
	var element []int = make([]int, 127)
	for i := 0; i < len(stream_string); i++ {
		if !((stream_string[i] >= 'a' && stream_string[i] <= 'z') || (stream_string[i] >= 'A' && stream_string[i] <= 'Z')) {
			return true
		}
		if element[stream_string[i]] > 0 {
			return true
		}
		element[stream_string[i]]++
	}
	return false
}

func Language_Check(arguments []string) bool {
	for i := 0; i < len(arguments); i++ {
		for j := 0; j < len(arguments[i]); j++ {
			if arguments[i][j] > 126 {
				error_handling.ErrorHandling("Character_Check", nil, true, "There is a character which is not allowed to input")
				return true
			}
		}
	}
	return false
}

func ArgumentHandling(arguments []string) (bool, string, string, string, string) {
	var flag_option bool = false
	var option string = ""
	if len(arguments) > 4 && error_handling.ErrorHandling("ArgumentHandling", nil, true, "There are too many arguments than it is supposed to be") {
		return true, "", "", "", ""
	} else if len(arguments) == 0 && error_handling.ErrorHandling("ArgumentHandling", nil, true, "Here supposed to be arguments.") {
		return true, "", "", "", ""
	}
	if Language_Check(arguments) {
		error_handling.ErrorHandling("ArgumentHandling", nil, true, "")
		return true, "", "", "", ""
	}

	flag_option, option = option_identification(arguments[0])

	if flag_option {
		if option == "reverse" {
			return false, option, "", "", ""
		}

		if len(arguments) == 1 {
			error_handling.ErrorHandling("ArgumentHandling", nil, true, "After option must be at least one argument")
			return true, "", "", "", ""
		} else if len(arguments) == 2 {
			return false, option, arguments[1], "", ""
		} else if len(arguments) == 3 {
			if !(len(arguments[1]) > 0 && len(arguments[2]) > 0) && option != "color" {
				error_handling.ErrorHandling("ArgumentHandling", nil, true, "A string or a banner must not be zero sized")
				return true, "", "", "", ""
			}
			if option == "color" {
				if !Check_letter(arguments[1]) {
					return false, "color", arguments[1], arguments[2], ""
				} else {
					if len(arguments[1]) > 0 {
						if banner_identification(arguments[2]) {
							return false, "color", "", arguments[1], arguments[2]
						}
					}
					error_handling.ErrorHandling("ArgumentHandling", nil, true, "The second argument must contain unique letters")
					return true, "", "", "", ""

				}
			}
			if !banner_identification(arguments[2]) && option != "color" {
				error_handling.ErrorHandling("ArgumentHandling", nil, true, "You have a problem with banner's position or you have too many strings than it is allowed")
				return true, "", "", "", ""
			}

			return false, option, arguments[1], arguments[2], ""

		} else if len(arguments) == 4 && option == "color" {
			if !(len(arguments[1]) > 0 && len(arguments[2]) > 0 && len(arguments[3]) > 0) {

				error_handling.ErrorHandling("ArgumentHandling", nil, true, "A string or a banner must not be zero sized")
				return true, "", "", "", ""
			}
			if !banner_identification(arguments[3]) {
				error_handling.ErrorHandling("ArgumentHandling", nil, true, "You have a problem with banner's position or you have too many strings than it is allowed")
				return true, "", "", "", ""
			}

			if Check_letter(arguments[1]) && error_handling.ErrorHandling("ArgumentHandling", nil, true, "At the second argument must be unique letters") {
				return true, "", "", "", ""
			}
			return false, "color", arguments[1], arguments[2], arguments[3]
		}
		error_handling.ErrorHandling("ArgumentHandling", nil, true, "Here too many argument than supposed to be")
		return true, "", "", "", ""

	} else {
		if len(arguments) == 2 {
			if !banner_identification(arguments[1]) {
				error_handling.ErrorHandling("ArgumentHandling", nil, true, "The string could not be as second argument")
			} else if len(arguments[0]) == 0 {
				error_handling.ErrorHandling("ArgumentHandling", nil, true, "The string size is could not be zero size")
			}
			return false, "", arguments[0], arguments[1], ""
		} else if len(arguments) == 1 {
			return false, "", arguments[0], "", ""
		}
	}

	error_handling.ErrorHandling("ArgumentHandling", nil, true, "Here is a problem of inputting format of option")
	return true, "", "", "", ""
}
