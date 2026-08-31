package main
import "fmt"
func increment(num int) {
	num++
    fmt.Println("Inside Function: ", num)
}
func main() {
	x := 5

    fmt.Println("Before Function Call: ", x)
    increment(x)
    fmt.Println("After Function Call: ", x)
}
