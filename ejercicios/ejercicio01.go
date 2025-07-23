package ejercicios

import (
	"strconv"
)

func ConvertirString(value string) (int, string) {
	numero, _ := strconv.Atoi(value)

	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}

}
