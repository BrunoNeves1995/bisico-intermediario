package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// CONCORRENCIA E  != DIFERENTE != DE PARALELISMO
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	go escrever("Ola mundo, no for while")
	go escrever("programando em go")
}

func escrever(texto string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
