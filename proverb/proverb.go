// Package proverb returns lines of "For Want of a Nail" proverb, equal to the amaunt of wort received
package proverb

// Proverb returns the poem
func Proverb(rhyme []string) []string {
	switch len(rhyme) {
	case 0:
		rhyme = []string{}
	case 1:
		rhyme = []string{"And all for the want of a " + rhyme[0] + "."}
	default:
		rhyme = fazRima(rhyme)
	}
	return rhyme
}

//FazRima cria o poema quando ele tem
func fazRima(rhyme []string) []string {
	retorno := []string{"For want of a " + rhyme[0] + " the " + rhyme[1] + " was lost."}
	size := len(rhyme) - 1
	for i := 1; i < size; i++ {
		linha := "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
		retorno = append(retorno, linha)
	}
	linha := "And all for the want of a " + rhyme[0] + "."
	retorno = append(retorno, linha)

	return retorno
}
