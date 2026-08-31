package main
import "fmt"
func double(num *int) {
	*num *= 2
}
func main() {
	x := 5
    fmt.Println("Before: ", x)
    double(&x)
    fmt.Println("After: ", x)
}
