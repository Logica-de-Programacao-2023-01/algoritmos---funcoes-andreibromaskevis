package main

import (
	"errors"
	"fmt"
)
func numerosPrimos(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("Número negativo")
	}
	primos := []int{}
	for i := 2; i <= n; i++ {
		if ehPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos, nil
}
func ehPrimo(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	num := 20
	primos, err := numerosPrimos(num)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(primos)
}
