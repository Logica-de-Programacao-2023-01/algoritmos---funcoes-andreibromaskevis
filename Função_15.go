package main

import (
	"errors"
	"fmt"
)
func numeroPresente(numero int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("Slice vazio")
	}
	for _, num := range slice {
		if num == numero {
			return true, nil
		}
	}
	return false, nil
}
func main() {
	slice := []int{1, 2, 3, 4, 5}
	numero := 3
	presente, err := numeroPresente(numero, slice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(presente)
}


