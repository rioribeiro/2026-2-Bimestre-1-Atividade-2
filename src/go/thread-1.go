package main

import (
    "fmt"       
    "sync"      
    "time"      
)

func minhaFuncao(wg *sync.WaitGroup) {
    defer wg.Done() 
    fmt.Println("Thread (goroutine) iniciada!")
    time.Sleep(2 * time.Second) 
    fmt.Println("Thread (goroutine) finalizada!")
}

func main() {
    var wg sync.WaitGroup 

    wg.Add(1)             
    go minhaFuncao(&wg)  

    wg.Wait()             
    fmt.Println("Programa principal finalizado!")
}
