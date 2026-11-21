package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Programa rodando")
}

func init() {
	fmt.Print("Iniciando as dependencias")
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}
	println()
}
