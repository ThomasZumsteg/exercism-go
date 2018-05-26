package protein

func FromCodon(codeon string) string {
	switch codeon {
	case "AUG":
		return "Methionine"
	case "UUU", "UUC":
		return "Phenylalanine"
	case "UUA", "UUG":
		return "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine"
	case "UAU", "UAC":
		return "Tyrosine"
	case "UGU", "UGC":
		return "Cysteine"
	case "UGG":
		return "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "STOP"
	}
	panic("Not a valid codon: " + codeon)
}

func FromRNA(rna string) (codons []string) {
	for i := 3; i <= len(rna); i += 3 {
		codon := FromCodon(rna[i-3 : i])
		if codon == "STOP" {
			return
		}
		codons = append(codons, codon)
	}
	return
}
