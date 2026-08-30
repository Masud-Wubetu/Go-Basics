package main

import (
    "fmt"
)

func main() {

    str := "Hello"

    for index, char := range str {
        fmt.Printf("Index : %d, Value: %c\n", index, char)
    }

}

