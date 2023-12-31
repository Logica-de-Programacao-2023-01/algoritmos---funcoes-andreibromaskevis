package main

import (
	"errors"
	"fmt"
)
func stringsComMaisDe5Caracteres(strings []string) ([]string, error) {
	if len(strings) == 0 {
		return nil, errors.New("Slice vazio")
	}
	resultado := []string{}
	for _, str := range strings {
		if len(str) > 5 {
			resultado = append(resultado, str)
		}
	}
	return resultado, nil
}
func main() {
	strings := []string{"Olá", "Mundo", "GPT-3", "OpenAI", "Go"}
	resultado, err := stringsComMaisDe5Caracteres(strings)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}




