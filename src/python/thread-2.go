package main

import (
    "fmt"
    "time"
)

func saudar(nome string, vezes int) {
    for i := 0; i < vezes; i++ {
            fmt.Printf("Olá, %s! (mensagem %d)\n", nome, i+1)
    }
}

func main() {
go saudar("Maria", 3)

time.Sleep(1 * time.Second)

}
