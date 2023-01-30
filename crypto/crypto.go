package crypto

func Cast(number int) string {
	var protein string
	switch number {
	case 0:
		protein = "A"
	case 1:
		protein = "C"
	case 2:
		protein = "G"
	case 3:
		protein = "T"
	case 4:
		protein = "U"
	}
	return protein
}
