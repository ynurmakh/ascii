package error_handling

import "fmt"

func ErrorHandling(function_name string, err error, consequence bool, message string) bool {
	if err != nil && !consequence {
		fmt.Print("Function name is: ", function_name)
		fmt.Println(". The Error message is: ", err)
		return true
	} else if err == nil && consequence {
		fmt.Print("Function name is: ", function_name)
		if len(message) > 0 {
			fmt.Print(",and it's error is: ")
			fmt.Println(" ", message)
		} else {
			fmt.Println()
		}
		return true
	}

	return false
}
