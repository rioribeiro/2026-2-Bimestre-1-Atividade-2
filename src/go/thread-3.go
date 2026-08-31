package main

import (
	"fmt"
	"sync"
	"time"
)

func trabalhador(numero int, tempoTrabalho time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Trabalhador %d começou\n", numero)
	time.Sleep(tempoTrabalho)
	fmt.Printf("Trabalhador %d terminou (levou %v)\n", numero, tempoTrabalho)
}

func main() {
	fmt.Println("Iniciando 5 trabalhadores...")
	inicio := time.Now()

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) 
		go trabalhador(i, 2*time.Second, &wg)
	}

	wg.Wait()

	tempoTotal := time.Since(inicio)

	fmt.Println("\nTodos os trabalhadores terminaram!")
	fmt.Printf("Tempo total: %.2fs\n", tempoTotal.Seconds())
	fmt.Println("(Se fosse sequencial, levaria ~10s)")
}