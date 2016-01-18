package dna

import (
	"fmt"
)

//Histogram counts the occurance of each nucleotide
type Histogram map[byte]int

//Dna is a string of nucleotides
type Dna map[byte]int

/*DNA creats a new dna strand.*/
func DNA(s string) Dna {
	d := Dna{'G': 0, 'T': 0, 'A': 0, 'C': 0}
	for _, n := range s {
		d[byte(n)]++
	}
	return d
}

/*Count counts the occurances of a nucleotide in a dna strand.*/
func (d Dna) Count(n byte) (int, error) {
	count, ok := d[n]
	if !ok {
		return 0, fmt.Errorf("%q is not a valid nucleotide", n)
	}
	return count, nil
}

/*Counts counts the occurance of all valid nucleotides in a dna strand.*/
func (d Dna) Counts() Histogram {
	return Histogram(d)
}
