package main

import (
    "fmt"
)

func main() {

    count := 1
   
    for {
        fmt.Println("count: ", count)
        count++

        if count > 5 {
            break
        }
    }
}

