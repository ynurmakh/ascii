package function

import (
	"fmt"

	"Adlet/main_function/container"
)

func Output(stream_string string, start, end, space int) {
	for j := 0; j < 8; j++ {
		for z := 0; z < space; z++ {
			fmt.Print(" ")
		}
		for i := start; i <= end; i++ {
			for q := 0; q < len(container.Format[stream_string[i]-32][j]); q++ {
				fmt.Print(string(container.Format[stream_string[i]-32][j][q]))
			}
		}
		for z := 0; z < space; z++ {
			fmt.Print(" ")
		}
		fmt.Println()

	}
}

func Center(terminal_width int, stream_string, banner string) {
	var line_size int = 0
	var count int = 0
	var end int = 0
	for i := 0; i < len(stream_string); i++ {
		for i+count < len(stream_string) && stream_string[i+count] != '\n' && line_size+container.Element_width[stream_string[i+count]-32] <= terminal_width {
			line_size += container.Element_width[stream_string[i+count]-32]
			count++
		}
		difference := terminal_width - line_size
		space := difference / 2
		end = i + count - 1
		Output(stream_string, i, end, space)
		if i+count < len(stream_string) && stream_string[i+count] == '\n' {
			end++
			fmt.Println()
		}
		count = 0
		line_size = 0
		i = end
	}
}
