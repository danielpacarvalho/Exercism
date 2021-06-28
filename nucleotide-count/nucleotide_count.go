// Package dna computes nucleotides
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = make(Histogram)
	h['A'] = 0
	h['C'] = 0
	h['T'] = 0
	h['G'] = 0

	for _, v := range d {
		valid := 0
		if v == 'A' {
			h['A']++
			valid++
		}
		if v == 'C' {
			h['C']++
			valid++
		}
		if v == 'G' {
			h['G']++
			valid++
		}
		if v == 'T' {
			h['T']++
			valid++
		}
		if valid == 0 {
			return h, errors.New("strand invalid")
		}
	}
	return h, nil
}
