package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gabrielgju/godesdecero/ejercicios"
)

var filename string = "./files/archivos/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.TablasMultiplicar()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumarTabla() {
	var texto string = ejercicios.TablasMultiplicar()
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo")
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

	archivo.Close()
}
