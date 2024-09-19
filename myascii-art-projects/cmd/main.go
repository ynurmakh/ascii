package main

import (
	"Adlet/argument"
	"Adlet/main_function/error_handling"
	"fmt"
	"os"

	color "Adlet/color/implementation"
	fs "Adlet/fs/implementation"
	justify "Adlet/justify/implementation"

	output "Adlet/output/implementation"
	reverse "Adlet/reverse/implementation"
)

func Question() {
	fmt.Println("\nHello there? ;)\n You supposed to have some problems with format of inputing data. In order to clarify you're problem, please answer me to some questions. :-)")

	fmt.Println("Do you have an option to use in your case? (y/n)")
	var answer string
	fmt.Scan(&answer)
	if answer == "y" {
		fmt.Println("So,What is your option ?(I just need the name of the option in lowercase)")
		fmt.Scan(&answer)
		if answer == "output" || answer == "justify" {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			if answer == "output" {
				fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			} else {
				fmt.Println("Example: go run . --align=right something standard")
			}
		} else if answer == "color" {
			fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		} else if answer == "reverse" {
			fmt.Println("Usage: go run . [OPTION]\nEX: go run . --reverse=<fileName>")
		} else if answer == "fs" {
			fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		} else {
			fmt.Println("There is no such option")
		}
	} else {
		if len(os.Args) > 3 {
			fmt.Println("There are many arguments than it is supposed to be")
		} else {
			fmt.Println("In this case, your banner must not be at the front of the  string")
		}
	}
}

func main() {
	mistake, option, stream_string, banner, color_stuff := argument.ArgumentHandling(os.Args[1:])
	if mistake && error_handling.ErrorHandling("main", nil, true, "") {
		Question()
		return
	}
	if len(stream_string) == 0 && len(option) != 0 && option != "reverse" && option != "color" && error_handling.ErrorHandling("main", nil, true, "The input string is zero size") {
		return
	}

	if option == "color" {
		if len(os.Args) == 4 {
			if len(banner) == 0 && error_handling.ErrorHandling("main", nil, true, "The input string is zero size") {
				return
			}
		}
	}
	if option == "output" {
		if len(banner) == 0 {
			if output.Solve(os.Args[1], stream_string, "standard.txt") && error_handling.ErrorHandling("main", nil, true, "") {
				return
			}
		} else {
			if output.Solve(os.Args[1], stream_string, banner+".txt") && error_handling.ErrorHandling("main", nil, true, "") {
				return
			}
		}
	} else if option == "reverse" { // Beforehand implementing the reverse option, you should first of all implement output file, in order to have something to read and convert to string
		if reverse.Solve(os.Args[1]) && error_handling.ErrorHandling("main", nil, true, "") {
			return
		}
	} else if option == "align" {
		if len(banner) == 0 {
			if justify.Solve(os.Args[1], stream_string, "standard") && error_handling.ErrorHandling("main", nil, true, "") {
				return
			}
		} else {
			if justify.Solve(os.Args[1], stream_string, banner) && error_handling.ErrorHandling("main", nil, true, "") {
				return
			}
		}
	} else if option == "color" {
		unique_string := stream_string
		input_string := banner
		if len(input_string) == 0 && len(color_stuff) == 0 {
			input_string = unique_string
		}
		if len(color_stuff) == 0 {
			if color.Solve(os.Args[1], unique_string, input_string, "standard") && error_handling.ErrorHandling("main", nil, true, "") {
				return
			}
		} else {
			if color.Solve(os.Args[1], unique_string, input_string, color_stuff) && error_handling.ErrorHandling("main", nil, true, "") {
				return
			}
		}
	} else {
		if len(banner) == 0 {
			if fs.Solve(stream_string, "standard.txt") && error_handling.ErrorHandling("main", nil, true, "") {
				return
			}
		} else {
			if fs.Solve(stream_string, banner+".txt") && error_handling.ErrorHandling("main", nil, true, "") {
				return
			}
		}
	}
}
