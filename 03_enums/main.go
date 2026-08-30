package main

import "fmt"

func getNumbers() (int, int) {
	return 5, 10
}

func main() {
   
    const A = "GFG"
    var B = "GeeksforGeeks"

    var helloWorld = A+ " " + B
    helloWorld += "!"
    fmt.Println(helloWorld) 

    
    const Correct = true
    fmt.Println("Go rules?", Correct)
}