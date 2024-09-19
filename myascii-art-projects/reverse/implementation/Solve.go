package implementation

import (
	"fmt"

	"Adlet/main_function/container"
	"Adlet/main_function/error_handling"
	Functionate "Adlet/main_function/implementation"
	"Adlet/reverse/support"
)

var banner_indicator string

func Generate(Mainstring_array [][]byte) string {
	var Sentence string
	var sum int = 0
	var space bool = true
	var previous_column bool = true
	var newline bool = false
	var count_space_column int = 0
	var element int = 0
	var stolbets int = 0

	double_quote := false
	for j := 0; j < len(Mainstring_array[0]); j++ {
		for i := 0; i < 8; i++ {
			if Mainstring_array[i][j] == '\n' {
				newline = true
				break
			}
			if Mainstring_array[i][j] != ' ' {
				space = false
			}
			element++
			sum += int(Mainstring_array[i][j]) * element * (stolbets + 1) * (i + 1)

		}

		stolbets++
		if space == true && previous_column == true {
			count_space_column++
		}

		if newline == true {

			Sentence += "\n"
			break
		}

		if (previous_column == false && space == true) || (j+1 == len(Mainstring_array[0]) && space == true && previous_column == true) || (previous_column == true && space == true && count_space_column == 6) {
			if sum == 118547 && banner_indicator == "shadow.txt" && stolbets == 3 {
				double_quote = true
				continue
			}
			if double_quote == true {
				double_quote = false
				continue
			}

			Sentence += string(byte(container.Hash_table[sum] + 32))
			sum = 0
			count_space_column = 0
			element = 0
			stolbets = 0
		}
		previous_column = space
		space = true
	}

	return Sentence
}

func Reading(stream_string, banner_name string) (bool, string) {
	var Sentence string = stream_string
	banner_indicator = banner_name
	var Answer string
	var Mainstring_array [][]byte
	var transmit string = ""

	for i := 0; i < len(Sentence); i++ { // Tranforming 1d array (Output_string) to 2d array
		if Sentence[i] != '\n' {
			transmit += string(Sentence[i])
		} else {
			transmit += "\n"
			Mainstring_array = append(Mainstring_array, []byte(transmit))
			transmit = ""
		}
	}

	for i := 0; i < len(Mainstring_array); i++ { // Cuttring 8xN sized array, and it has been passed to SOlve function, in order to get a string of representation of MainString.
		if len(Mainstring_array[i]) > 1 {
			Answer += Generate(Mainstring_array[i : i+8])
			i += 7
			continue
		} else if len(Mainstring_array[i]) == 1 {
			Answer += "\n"
		}
	}

	return false, Answer
}

func GetName(name string) (bool, string) {
	if len(name) <= 4 {
		return true, ""
	}
	if name[len(name)-4:] != ".txt" {
		return true, ""
	}
	return false, name[10:]
}

func Solve(first string) bool {
	flag, file_name := GetName(first)
	if flag && error_handling.ErrorHandling("Solve", nil, true, "Ther is problem with file name, and the file's extension should be \".txt\"") {
		return true
	}
	flag, banner_name := support.Reading_file("banner.txt")
	if flag && error_handling.ErrorHandling("Solve", nil, true, "") {
		return true
	}
	Functionate.Obrabotka(banner_name)
	flag, stream_string := support.Reading_file(file_name)
	if flag && error_handling.ErrorHandling("Solve", nil, true, "") {
		return true
	}

	flag, Answer := Reading(stream_string, banner_name)
	if flag && error_handling.ErrorHandling("Solve", nil, true, "") {
		return true
	}
	fmt.Print(Answer)

	return false
}
