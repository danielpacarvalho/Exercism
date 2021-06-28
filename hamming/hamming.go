package hamming

import "errors"

//Calculate Hamming distance
func Distance(a, b string) (int, error) {
	var hammingDistance int = 0
	//Distance can only compare strands of the same size
	if len(a) != len(b) {
		return 0, errors.New("strands of differente length")
	}

	//Calculate the Hamming Distance
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
