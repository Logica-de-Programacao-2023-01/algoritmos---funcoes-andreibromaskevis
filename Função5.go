package main

import "fmt"
func encontrarPosicao(slice []int, valor int) int {
	for i, num := range slice {
		if num == valor {
			return i
		}
	}
	return -1
}
func main() {
	slice := []int{10, 5, 8, 12, 3}
	valor := 8
	posicao := encontrarPosicao(slice, valor)
	fmt.Println(posicao)
}
