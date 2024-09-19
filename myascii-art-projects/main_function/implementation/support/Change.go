package support

func Change(Mainstring string) string {
	var sentence string
	for i := 0; i < len(Mainstring); i++ {
		if Mainstring[i] != '\r' {
			sentence += string(Mainstring[i])
		}
	}
	return sentence
}
