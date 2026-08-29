package main

import "fmt"

func getNumbers() (int, int) {
	return 5, 10
}

const PI = 3.14

func main() {
    const GFG = "GeeksforGeeks"
    fmt.Println("Hello", GFG)

    fmt.Println("Happy", PI, "Day")

    const Correct = true
    fmt.Println("Go rules?", Correct)
}