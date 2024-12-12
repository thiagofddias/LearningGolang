package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	_, err := os.Open("test.txt")
	if err != nil {
		panic("Erro ao abrir o arquivo")
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado com sucesso")
		}
	}()

	ReadFile()

	defer fmt.Println("Finalizando a manipulacao do arquivo")
}

// func ShowText() {
// 	fmt.Println("Finalizando a manipualacao do arquivo")
// }

// func main() {
// 	file, err := os.Create("test.txt")
// 	defer file.Close()
// 	defer ShowText()

// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err = file.Write([]byte("Teste"))

// 	if err != nil {
// 		panic(err)
// 	}
// }
