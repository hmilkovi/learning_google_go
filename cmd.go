package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))

    // Exercise 1.1
    fmt.Printf("Command: %s\n", os.Args[0])

    // Exercise 1.2
    for i, cmd := range(os.Args[1:]) {
        fmt.Printf("%d:%s\n", i, cmd)
    }

    // Exercise 1.3  TO DO banchmarks
}