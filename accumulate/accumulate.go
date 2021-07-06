package accumulate

import "strings"

func Accumulate(lista []string, comando string) []string {
	retorno := make([]string, len(lista))
	if strings.Compare(comando, "echo") == 0 {
		retorno = append(retorno, lista...)
	}
	if strings.Compare(comando, "capitalize") == 0 {
		for _, palavra := range lista {
			retorno = append(retorno, strings.ToUpper(palavra))
		}
	}
	if strings.Compare(comando, "strings.ToUpper") == 0 {
		for _, palavra := range lista {
			retorno = append(retorno, strings.ToUpper(palavra))
		}
	}
	return retorno
}
