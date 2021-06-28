// Package strand transcripts a DNA sequence to a RNA sequence
package strand

//ToRNA converts a dna strand into a rna strand, it returns "" if the strand received is empty
func ToRNA(d string) string {
	dna := []byte(d)
	var rna []byte

	for _, v := range dna {
		if v == 'G' {
			rna = append(rna, 'C')
		}
		if v == 'C' {
			rna = append(rna, 'G')
		}
		if v == 'T' {
			rna = append(rna, 'A')
		}
		if v == 'A' {
			rna = append(rna, 'U')
		}
	}
	return string(rna)
}
