package main

import (
	"errors"
	"fmt"
	"strings"
)
func substituirVogaisPorAsterisco(str string) (string, error) {
	if str == "" {
		return "", errors.New("String vazia")
	}
	vogais := "aeiouAEIOU"
	resultado := ""
	for _, char := range str {
		if strings.ContainsRune(vogais, char) {
			resultado += "*"
		} else {
			resultado += string(char)
		}
	}
	return resultado, nil
}
func main() {
	str := "Hello, World!"
	resultado, err := substituirVogaisPorAsterisco(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}




