package main

import (
    "fmt"
    "log"

    "rsc.io/quote"

    "go-experiments/greetings"
)

func main() {
    log.SetPrefix("greetings: ") // set logger prefix
    //log.SetFlags(0) // disable printing time, source file, line number info

    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())

    fmt.Print("\nTesing with loop\n\n")
    names := []string {"Amy", "Chris", "Elizabeth"}
    for _, name := range names {
        message, err := greetings.Hello(name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(message)
    }

    fmt.Print("\nTesing with slice\n\n")
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // no error, print the returned map
    fmt.Println(messages)
}
