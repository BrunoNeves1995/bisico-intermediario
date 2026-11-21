package main

import (
	"fmt"
	"time"
)

func main() {
	var nome string
	println("Executando antes da gorountine iniciar")
	go func() {
		nome = "sms 2"
		time.Sleep(1 * time.Second)
		fmt.Println("Programa rodando 1")
		println(nome)
	}()
	go func() {
		nome = "email 2"
		time.Sleep(3 * time.Second)
		fmt.Println("Programa rodando 2")
		println(nome)
	}()
	go func() {
		nome = "push notification 3"
		time.Sleep(2 * time.Second)
		fmt.Println("Programa rodando 3")
		println(nome)
	}()

	println("Executando depois da gorountine iniciar")
	select {}

}
