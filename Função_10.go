package main

import (
	"errors"
	"fmt"
	"sort"
)
func ordenarValores(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}
	sortedSlice := make([]int, len(slice))
	copy(sortedSlice, slice)
	sort.Ints(sortedSlice)

	return sortedSlice, nil
}
func main() {
	valores := []int{7, 3, 9, 2, 1, 5}
	resultado, err := ordenarValores(valores)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultado)
}

