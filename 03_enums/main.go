package main

import (
    "fmt"
)

func main() {
   
    a := 4
   
    b := &a 
    fmt.Println(*b) 
    *b = 7 
    fmt.Println(a) 
}