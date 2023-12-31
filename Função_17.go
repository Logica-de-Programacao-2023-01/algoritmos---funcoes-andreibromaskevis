package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)
func ordenarStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}
	sort.Strings(slice)
	resultado := strings.Join(slice, ", ")
	return resultado, nil
}
func main() {
	slice := []string{"banana", "abacaxi", "laranja", "uva"}
	resultado, err := ordenarStrings(slice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}


