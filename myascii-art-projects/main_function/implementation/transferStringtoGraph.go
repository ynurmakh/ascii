package implementation

import (
	"Adlet/main_function/container"
)

var sentence string

func operate(start_index, end_index int, Format [][]string) string {
	if start_index == end_index {
		return ""
	}
	var object string
	for i := 0; i < 8; i++ {
		for j := start_index; j < end_index; j++ {
			object += string(Format[sentence[j]-32][i])
		}
		if i != 7 {
			object += "\n"
		}
	}
	return object
}

func TransferStringToGraph(string_stream string) {
	sentence = string_stream
	var start, end int = 0, 0
	var flag bool = false
	for i := 0; i < len(string_stream); i++ {
		if string_stream[i] == '\n' {
			end = i
			container.Result += operate(start, end, container.Format)
			start = i + 1
			container.Result += "\n"
		} else {
			flag = true
		}
	}
	container.Result += operate(start, len(string_stream), container.Format)
	if flag == true {
		container.Result += "\n"
	}
}
