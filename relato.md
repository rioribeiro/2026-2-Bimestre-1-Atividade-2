# Relatório de implementação de linha de execução em Go

## Introdução

Este relato faz parte do processo avaliativo da disciplina de sistemas operacionas no curso superior em análise e desenvolvimento de sistemas, ofertado na Diretoria acadêmica de gestão e tecnologia da informação no campus natal-central do instituto federal de educação, ciência e tecnologia do rio grande do norte.

Tem como objetivo principal relatar como implementar linhas de execução na linguagem Go.

O grupo de trabalho foi formado por João Victor, Rio Ribeiro e Wherverton Cruz.

## Implementando múltiplas linhas de execução em Go

### Informações gerais sobre Go

A linguagem Go, também chamada de Golang, foi criada pelo Google em 2009.

- Paradigma: imperativa, concorrente e fortemente tipada.

- Objetivo: oferecer simplicidade, eficiência e suporte nativo à concorrência por meio das goroutines e canais.

- Disponibilidade: é open source e pode ser obtida gratuitamente no site oficial golang.org.

### Criando linhas de execução

Em Go, uma linha de execução é chamada de goroutine. Para criar uma goroutine, basta usar a palavra-chave go antes da chamada de função.
Exemplo adaptado do thread-1.py:

```
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
```

### Passando valores para linhas de execução

As goroutines podem receber parâmetros normalmente, como qualquer função.
Exemplo adaptado do thread-2.py:
```
func saudar(nome string, vezes int) {
    for i := 0; i < vezes; i++ {
        fmt.Printf("Olá, %s! (mensagem %d)\n", nome, i+1)
    }
}

func main() {
    go saudar("Maria", 3)
    time.Sleep(1 * time.Second)
}
```

### Múltiplas linhas de execução

Para executar várias goroutines simultaneamente, usamos sync.WaitGroup para sincronizar.
Exemplo adaptado do thread-3.py:
```
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
```

## Considerações finais

A implementação em Go demonstrou como as goroutines tornam a concorrência simples e eficiente. Comparando com Python, onde usamos a biblioteca threading, em Go o suporte é nativo e integrado à linguagem. O uso de sync.WaitGroup facilita a sincronização, e o modelo de goroutines permite criar milhares de execuções concorrentes com baixo custo.

Portanto, o trabalho evidenciou a importância da concorrência em sistemas operacionais e mostrou como diferentes linguagens oferecem mecanismos distintos para lidar com múltiplas linhas de execução.
