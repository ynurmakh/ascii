package implementation

import (
	"fmt"

	"Adlet/ascii-art/support"
	"Adlet/main_function/container"
	"Adlet/main_function/error_handling"
	main_function "Adlet/main_function/implementation"
	main_function_support "Adlet/main_function/implementation/support"

	"github.com/fatih/color"
)

var (
	Element_flag []bool = make([]bool, 256)
	colour       string
)

func GiveMeColour(character string) string {
	switch colour {
	case "red":
		return color.New(color.FgRed).SprintfFunc()(character)
	case "green":
		return color.New(color.FgGreen).SprintfFunc()(character)
	case "blue":
		return color.New(color.FgBlue).SprintfFunc()(character)
	case "black":
		return color.New(color.FgBlack).SprintfFunc()(character)
	case "white":
		return color.New(color.FgWhite).SprintfFunc()(character)
	case "yellow":
		return color.New(color.FgYellow).SprintfFunc()(character)
	case "purple":
		return color.New(color.FgHiMagenta).SprintfFunc()(character)
	}

	return "NO"
}

func set_target(target string) {
	if len(target) != 0 {
		for i := 0; i < len(target); i++ {
			Element_flag[target[i]] = true
		}
	} else {
		for i := 32; i <= 126; i++ {
			Element_flag[i] = true
		}
	}
}

func Output(stream_string string, start, end int) {
	for j := 0; j < 8; j++ {

		for i := start; i <= end; i++ {
			for q := 0; q < len(container.Format[stream_string[i]-32][j]); q++ {
				if Element_flag[stream_string[i]] == true {
					fmt.Print(GiveMeColour(string(container.Format[stream_string[i]-32][j][q])))
				} else {
					fmt.Print(string(container.Format[stream_string[i]-32][j][q]))
				}
			}
		}

		fmt.Println()

	}
}

func Color(terminal_width int, stream_string, banner string) {
	var line_size int = 0
	var count int = 0
	var end int = 0
	for i := 0; i < len(stream_string); i++ {
		for i+count < len(stream_string) && stream_string[i+count] != '\n' && line_size+container.Element_width[stream_string[i+count]-32] <= terminal_width {
			line_size += container.Element_width[stream_string[i+count]-32]
			count++
		}
		end = i + count - 1
		Output(stream_string, i, end)
		if i+count < len(stream_string) && stream_string[i+count] == '\n' {
			end++
			fmt.Println()
		}
		count = 0
		line_size = 0
		i = end
	}
}

func GetnameColor(name string) string {
	var object string = ""
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] != '=' {
			object = string(name[i]) + object
		} else {
			return object
		}
	}
	return ""
}

func CheckColor(name string) bool {
	if !(name == "red" || name == "green" || name == "blue" || name == "white" || name == "black" || name == "purple" || name == "yellow") {
		return true
	}
	return false
}

func Solve(color, target, stream_string, banner string) bool {
	colour = GetnameColor(color)
	if CheckColor(colour) && error_handling.ErrorHandling("Solve", nil, true, "The name of the color is not correct") {
		return true
	}
	set_target(target)
	flag, terminal_width := support.Find_Terminal_Width()
	if flag && error_handling.ErrorHandling("Solve", nil, true, "") {
		return true
	}
	if main_function.Main_Operation(stream_string, banner+".txt", false) && error_handling.ErrorHandling("Solve", nil, true, "") {
		return true
	}
	stream_string = main_function_support.Modification(stream_string)
	Color(terminal_width, stream_string, banner)

	return false
}
