package main

import "fmt"

import "rsc.io/quote"

import "github.com/amybabay/go-experiments/internal/greetings"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())

    message := greetings.Hello("Amy")
    fmt.Println(message)
}
