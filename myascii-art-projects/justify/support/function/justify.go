package function

import (
	"fmt"

	"Adlet/main_function/container"
)

var (
	ignore1, ignore2 int
	flag             bool = false
)

func outputjustify(stream_string string, start, end, space int) {
	for j := 0; j < 8; j++ {

		for i := start; i <= end; i++ {
			if !flag {
				if stream_string[i] != ' ' {
					for q := 0; q < len(container.Format[stream_string[i]-32][j]); q++ {
						fmt.Print(string(container.Format[stream_string[i]-32][j][q]))
					}
				} else {
					if i > ignore1 && i < ignore2 {
						for z := 0; z < space; z++ {
							fmt.Print(" ")
						}
					}
				}
			} else {
				for q := 0; q < len(container.Format[stream_string[i]-32][j]); q++ {
					fmt.Print(string(container.Format[stream_string[i]-32][j][q]))
				}
			}
		}
		fmt.Println()

	}
}

func Amount(stream_string string, start, end int) int {
	var word bool = false
	var count int = 0
	var once bool = true
	var change_once bool = true
	for i := start; i <= end; i++ {
		if stream_string[i] != ' ' {
			word = true
			if once {
				once = false
				ignore1 = i
			}
			if change_once {
				change_once = false
				ignore2 = i
			}
		} else if word {
			count++
			word = false
			change_once = true
		}
	}
	if word {
		count++
	}
	return count
}

func Justify(terminal_width int, stream_string, banner string) {
	var line_size int = 0
	var count int = 0
	var end int = 0
	var duplicate_space int = 0
	var save int
	var space int
	var switch_to bool = false
	for i := 0; i < len(stream_string); i++ {
		save = count
		for i+count < len(stream_string) && stream_string[i+count] != '\n' && duplicate_space+container.Element_width[stream_string[i+count]-32] <= terminal_width {
			if stream_string[i+count] != ' ' {
				duplicate_space += container.Element_width[stream_string[i+count]-32]
			}
			count++
		}

		difference := terminal_width - duplicate_space
		word_amount := Amount(stream_string, i, i+count-1)
		if word_amount == 0 || word_amount == 1 {
			space = 0
		} else {
			space = difference / (word_amount - 1)
			if space < 6 {
				switch_to = true
			}
		}

		if switch_to {
			count = save
			for i+count < len(stream_string) && stream_string[i+count] != '\n' && line_size+container.Element_width[stream_string[i+count]-32] <= terminal_width {

				line_size += container.Element_width[stream_string[i+count]-32]
				count++
			}
			space = 0
			flag = true
			switch_to = false
		}

		end = i + count - 1
		outputjustify(stream_string, i, end, space)
		if i+count < len(stream_string) && stream_string[i+count] == '\n' {
			end++
			fmt.Println()
		}

		count = 0
		line_size = 0
		duplicate_space = 0
		flag = false
		i = end
	}
}
