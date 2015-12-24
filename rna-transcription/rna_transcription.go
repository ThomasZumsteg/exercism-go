package strand

//TestVersion is the unit tests that these functions will pass
const TestVersion = 1

//dnaToRna maps a dna nucleotide to an rna nucleotide
var dnaToRna = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

/*ToRna converts a string of dna to rna*/
func ToRna(dnaStrand string) string {
	rnaStrand := ""
	for _, dna := range dnaStrand {
		//Error not tested for and ignored
		rna, _ := dnaToRna[dna]
		rnaStrand += rna
	}
	return rnaStrand
}
