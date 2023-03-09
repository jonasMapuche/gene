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

func Codon(triple string) []string {
	var size int = len(triple)
	var amino []string
	var text = []rune(triple)
	for count := 1; count <= size; count++ {
		if (count % 3) == 0 {
			var code = text[count-2] + text[count-1] + text[count]
			switch string(code) {
			case "UUU":
				amino = append(amino, "Fenilalanina")
			case "UUC":
				amino = append(amino, "Fenilalanina")
			case "UUA":
				amino = append(amino, "Leucina")
			case "UUG":
				amino = append(amino, "Leucina")
			case "CUU":
				amino = append(amino, "Leucina")
			case "CUC":
				amino = append(amino, "Leucina")
			case "CUA":
				amino = append(amino, "Leucina")
			case "CUG":
				amino = append(amino, "Leucina")
			case "AUU":
				amino = append(amino, "Isoleucina")
			case "AUC":
				amino = append(amino, "Leucina")
			case "AUA":
				amino = append(amino, "Leucina")
			case "AUG":
				amino = append(amino, "Metionina [Start]")
			case "GUU":
				amino = append(amino, "Valina")
			case "GUC":
				amino = append(amino, "Valina")
			case "GUA":
				amino = append(amino, "Valina")
			case "GUG":
				amino = append(amino, "Valina")
			case "UCU":
				amino = append(amino, "Serina")
			case "UCC":
				amino = append(amino, "Serina")
			case "UCA":
				amino = append(amino, "Serina")
			case "UCG":
				amino = append(amino, "Serina")
			case "CCU":
				amino = append(amino, "Prolina")
			case "CCC":
				amino = append(amino, "Prolina")
			case "CCA":
				amino = append(amino, "Prolina")
			case "CCG":
				amino = append(amino, "Prolina")
			case "ACU":
				amino = append(amino, "Treonina")
			case "ACC":
				amino = append(amino, "Treonina")
			case "ACA":
				amino = append(amino, "Treonina")
			case "ACG":
				amino = append(amino, "Treonina")
			case "GCU":
				amino = append(amino, "Alanina")
			case "GCC":
				amino = append(amino, "Alanina")
			case "GCA":
				amino = append(amino, "Alanina")
			case "GCG":
				amino = append(amino, "Alanina")
			case "UAU":
				amino = append(amino, "Tirosina")
			case "UAC":
				amino = append(amino, "Tirosina")
			case "UAA":
				amino = append(amino, "[Stop]")
			case "UAG":
				amino = append(amino, "[Stop]")
			case "CAU":
				amino = append(amino, "Histidina")
			case "CAC":
				amino = append(amino, "Histidina")
			case "CAA":
				amino = append(amino, "Glutamina")
			case "CAG":
				amino = append(amino, "Glutamina")
			case "AAU":
				amino = append(amino, "Asparagina")
			case "AAC":
				amino = append(amino, "Asparagina")
			case "AAA":
				amino = append(amino, "Lisina")
			case "GAU":
				amino = append(amino, "Ácido Aspártico")
			case "GAC":
				amino = append(amino, "Ácido Aspártico")
			case "GAA":
				amino = append(amino, "Ácido Glutâmico")
			case "GAG":
				amino = append(amino, "Ácido Glutâmico")
			case "UGU":
				amino = append(amino, "Cysteine")
			case "UGC":
				amino = append(amino, "Cysteine")
			case "UGA":
				amino = append(amino, "[Stop]")
			case "UGG":
				amino = append(amino, "Tryptophan")
			case "CGU":
				amino = append(amino, "Arginina")
			case "CGC":
				amino = append(amino, "Arginina")
			case "CGA":
				amino = append(amino, "Arginina")
			case "CGG":
				amino = append(amino, "Arginina")
			case "AGU":
				amino = append(amino, "Serina")
			case "AGC":
				amino = append(amino, "Serina")
			case "AGA":
				amino = append(amino, "Arginina")
			case "AGG":
				amino = append(amino, "Arginina")
			case "GGU":
				amino = append(amino, "Glicina")
			case "GGC":
				amino = append(amino, "Glicina")
			case "GGA":
				amino = append(amino, "Glicina")
			case "GGG":
				amino = append(amino, "Glicina")
			}
		}
	}
	return amino
}
