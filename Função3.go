package main

import (
	"fmt"
	"strings"
)
func concatenarStrings(strs []string) string {
	return strings.Join(strs, "")
}
func main() {
	slice := []string{"Olá", ", ", "mundo!"}
	concatenacao := concatenarStrings(slice)
	fmt.Println(concatenacao)
}
