package main

import (
	"fmt"

	"github.com/gabrielgju/godesdecero/ejercicios"
)

func main() {
	num, texto := ejercicios.ConvertirString("101")

	fmt.Println(num)
	fmt.Println(texto)
}
