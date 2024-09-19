package implementation

import (
	"os"

	"Adlet/main_function/container"
	"Adlet/main_function/error_handling"
	"Adlet/main_function/implementation/support"
)

func Reading_file(file_name string) (bool, string) {
	file_name = "../main_function/banner/" + file_name
	file, err := os.Open(file_name)
	if flag := error_handling.ErrorHandling("Reading_file", err, false, ""); flag == true {
		return true, ""
	}
	defer file.Close()

	fileInfo, err := os.Stat(file_name)
	if flag := error_handling.ErrorHandling("Reading_file", err, false, ""); flag == true {
		return true, ""
	}

	data := make([]byte, fileInfo.Size())
	n, err := file.Read(data)
	if flag := error_handling.ErrorHandling("Reading_file", err, false, ""); flag == true {
		return true, ""
	}

	content := string(data[:n])
	if file_name == "../main_function/banner/thinkertoy.txt" {
		content = support.Change(content)
	}

	return false, content
}

func convert(content string) {
	var array_sentence [][]string = make([][]string, 126-31)
	var array_sentence1d []string
	var object1d string
	var object string
	var flag bool = false
	var count_row int = 0
	var sum int = 0

	var elements int = 1
	var stroka int = 0
	var stolbets int = 0
	for i := 0; i < len(content); i++ {
		if content[i] >= 32 && content[i] <= 126 {
			flag = true
			object += string(content[i])
			object1d += string(content[i])
			stolbets++
			sum += int(content[i]) * ((stroka + 1) * (stolbets)) * (elements) // Here for outputint data
			elements += 8
		} else if flag {
			flag = false
			stolbets = 0
			stroka++
			elements = stroka + 1
			array_sentence[count_row] = append(array_sentence[count_row], object)
			object = ""
			object1d += string(content[i])

			if i+1 == len(content) || content[i+1] == '\n' {
				array_sentence1d = append(array_sentence1d, object1d)
				// In order to have unique index to the elements
				if int64(sum) > container.Max { // Here for outputint data
					container.Max = int64(sum)
				}
				container.Hash_table[sum] = byte(count_row) // Here for outputint data
				// fmt.Println(sum, "--->", Container.Sum_to_index_of_Format_array[sum])
				count_row++
				i++
				object1d = ""
				sum = 0 // Here for outputing data
				elements = 1
				stroka = 0
			}
		}
	}
	container.Format = array_sentence
	container.Format1d = array_sentence1d
}

func find_width(Format [][]string) {
	var Width []int = make([]int, len(Format))

	for i := 0; i < len(Format); i++ {
		Width[i] = len(Format[i][0])
	}

	container.Element_width = Width
}

func Obrabotka(file_name string) bool {
	var content string
	var flag bool = false
	if flag, content = Reading_file(file_name); flag == true {
		return true
	}

	convert(content)
	find_width(container.Format)

	return false
}
