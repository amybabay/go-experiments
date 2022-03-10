package main

import (
    "fmt"

    "rsc.io/quote"

    "github.com/amybabay/go-experiments/internal/greetings"
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())

    message := greetings.Hello("Amy")
    fmt.Println(message)
}
