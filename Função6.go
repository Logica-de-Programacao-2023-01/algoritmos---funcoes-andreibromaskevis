package main

import (
	"errors"
	"fmt"
	"strings"
)
func concatenarS2(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}
	resultado := strings.Join(slice, ",")
	return resultado, nil
}
func main() {
	slice := []string{"Olá", "Mundo", "Go"}
	concatenado, err := concatenarS2(slice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(concatenado)
}


