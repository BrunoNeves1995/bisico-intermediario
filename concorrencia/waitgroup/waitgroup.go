package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	fmt.Println("rodando primeira go func() -> escrever")
	go func() {
		escrever("Ola mundo, no for while")
		waitgroup.Done() // retira uma go routine do contador waitgroup.Add(1)
	}()
	
	fmt.Println("rodando segunda go func() -> escrever")
	go func() {
		go escrever("programando em go")
		waitgroup.Done() // retira uma go routine do contador waitgroup.Add(0)
	}()

	waitgroup.Wait() // espera a contagem da go routine chegar no zero waitgroup.Add(0)
	fmt.Print("As goroutine terminar terminou a execucao")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
