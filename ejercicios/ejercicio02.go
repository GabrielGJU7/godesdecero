package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablasMultiplicar() string {

	scanner := bufio.NewScanner(os.Stdin)

	var num int
	var err error
	var texto string

	for {
		fmt.Println("Ingresa un numero: ")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			}
			break
		}
	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", num, i, num*i)
	}
	return texto
}
