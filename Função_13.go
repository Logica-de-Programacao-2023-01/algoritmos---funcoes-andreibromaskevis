package main

import (
	"errors"
	"fmt"
	"strconv"
)
func somaDigitos(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("Número negativo")
	}
	numeroStr := strconv.Itoa(numero)
	soma := 0
	for _, digitoChar := range numeroStr {
		digito, err := strconv.Atoi(string(digitoChar))
		if err != nil {
			return 0, err
		}
		soma += digito
	}
	return soma, nil
}
func main() {
	numero := 12345
	resultado, err := somaDigitos(numero)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}


