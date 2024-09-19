package implementation

import (
	"Adlet/main_function/container"
)

func Output(terminal_width int, stream_string string) {
	var (
		Format        []string = container.Format1d
		Element_width []int    = container.Element_width
	)

	var (
		line_width                       int = 0
		count                            int = 0
		result                           string
		save_index_newline_in_Mainstring int  = -1
		overflow                         bool = false
		istherecharacter                 bool = false
	)

	for i := 0; i < len(stream_string); i++ {

		for line_width <= terminal_width && i+count < len(stream_string) { /* Counting the quantity of elements which could be printed to standard output in a line.*/
			if stream_string[i+count] == '\n' {

				save_index_newline_in_Mainstring = i + count

				break
			}
			if line_width+Element_width[stream_string[i+count]-32] <= terminal_width {
				line_width += Element_width[stream_string[i+count]-32]
				count++
			} else {
				overflow = true
				break
			}
		}
		// fmt.Println("The length of the elements which are staying at current line:", line_width)

		if count > 0 {

			line_width = 0
			var count_depth int = 0

			for height := 0; height < 8; height++ {
				for q := 0; q < count; q++ {
					for each := 0; each < len(Format[stream_string[i+q]-32]); each++ {
						if Format[stream_string[i+q]-32][each] == '\n' {
							count_depth++
							if count_depth > height {
								break
							}
						} else if count_depth == height {
							istherecharacter = true
							result += string(Format[stream_string[i+q]-32][each])
						}
					}
					count_depth = 0
				}
				if height+1 != 8 {
					result += "\n"
				}
			}
		}
		if overflow {
			overflow = false

			result += "\n"
		}
		if i == save_index_newline_in_Mainstring {
			result += "\n"
			i++
		}

		i += count
		i--
		count = 0
	}
	if istherecharacter {
		result += "\n"
	}

	container.Result = result
}
