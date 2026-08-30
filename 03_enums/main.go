package main

import (
    "fmt"
)

func main() {

    sum := 0
    num := 1

    for num <= 5 {
        sum += num
        num++
    }

    fmt.Println("Sum: ", sum)

}

