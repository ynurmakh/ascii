package support

func Modification(Sentence string) string { // Modifies the escape sequence from the original string line into desired characters.
	for i := 0; i < len(Sentence); i++ {
		if i+1 < len(Sentence) {
			if Sentence[i] == '\t' {
				Sentence = Sentence[:i] + "\\" + "t" + Sentence[i+1:]
			}
		} else {
			if Sentence[i] == '\t' {
				Sentence = Sentence[:i] + "\\" + "t"
			}
		}
	}
	for i := 0; i < len(Sentence); i++ {
		if i+1 < len(Sentence) {
			if Sentence[i] == '\r' {
				Sentence = Sentence[:i] + "\\" + "r" + Sentence[i+1:]
			}
		} else {
			if Sentence[i] == '\r' {
				Sentence = Sentence[:i] + "\\" + "r"
			}
		}
	}
	for i := 0; i < len(Sentence); i++ {
		if i+1 < len(Sentence) {
			if Sentence[i] == '\a' {
				Sentence = Sentence[:i] + "\\" + "a" + Sentence[i+1:]
			}
		} else {
			if Sentence[i] == '\a' {
				Sentence = Sentence[:i] + "\\" + "a"
			}
		}
	}
	for i := 0; i < len(Sentence); i++ {
		if i+1 < len(Sentence) {
			if Sentence[i] == '\b' {
				Sentence = Sentence[:i] + "\\" + "b" + Sentence[i+1:]
			}
		} else {
			if Sentence[i] == '\b' {
				Sentence = Sentence[:i] + "\\" + "b"
			}
		}
	}
	for i := 0; i < len(Sentence); i++ {
		if i >= 0 && i+1 < len(Sentence) && Sentence[i] == '\\' && Sentence[i+1] == 't' {
			if i+2 == len(Sentence) {
				Sentence = Sentence[:i] + "    "
			} else {
				Sentence = Sentence[:i] + "    " + Sentence[i+2:]
			}
		}
		if i >= 0 && i+1 < len(Sentence) && Sentence[i] == '\\' && Sentence[i+1] == 'a' {
			if i+2 == len(Sentence) {
				Sentence = Sentence[:i]
			} else {
				Sentence = Sentence[:i] + Sentence[i+2:]
				i--
			}
		}
		if i >= 0 && i+1 < len(Sentence) && Sentence[i] == '\\' && Sentence[i+1] == 'b' {
			if i+2 == len(Sentence) {
				if i == 0 {
					Sentence = ""
				} else {
					Sentence = Sentence[:i-1]
				}
			} else {
				if i == 0 {
					Sentence = Sentence[i+2:]
				} else {
					Sentence = Sentence[:i-1] + Sentence[i+2:]
					i -= 2
				}
			}
		}
		if i >= 0 && i+1 < len(Sentence) && Sentence[i] == '\\' && Sentence[i+1] == 'r' {
			if i == 0 {
				if i+2 == len(Sentence) {
					Sentence = Sentence[:i]
				} else {
					Sentence = Sentence[i+2:]
				}
			} else {
				if i+2 == len(Sentence) {
					Sentence = Sentence[:i]
				} else {
					remain_size := len(Sentence[i+2:])
					if remain_size >= i {
						Sentence = Sentence[i+2:]
					} else {
						Sentence = Sentence[i+2:] + Sentence[remain_size:i]
					}
				}
			}
			i = 0
		}
		if i >= 0 && i+1 < len(Sentence) && Sentence[i] == '\\' && Sentence[i+1] == 'n' {
			if i+2 == len(Sentence) {
				Sentence = Sentence[:i] + "\n"
			} else {
				Sentence = Sentence[:i] + "\n" + Sentence[i+2:]
			}
		}
	}
	return Sentence
}
