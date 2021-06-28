// Package scrabble is used to eletronically play the board game scrabble
package scrabble

import "strings"

// Score returns the point value of a given word
func Score(palavra string) int {
	palavra = strings.ToUpper(palavra)
	retorno := 0
	p := []byte(palavra)
	for _, v := range p {
		if v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' || v == 'L' || v == 'N' || v == 'R' || v == 'S' || v == 'T' {
			retorno++
		}
		if v == 'D' || v == 'G' {
			retorno += 2
		}
		if v == 'B' || v == 'C' || v == 'M' || v == 'P' {
			retorno += 3
		}
		if v == 'F' || v == 'H' || v == 'V' || v == 'W' || v == 'Y' {
			retorno += 4
		}
		if v == 'K' {
			retorno += 5
		}
		if v == 'J' || v == 'X' {
			retorno += 8
		}
		if v == 'Q' || v == 'Z' {
			retorno += 10
		}
	}
	return retorno
}
