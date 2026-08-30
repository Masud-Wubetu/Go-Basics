package main
import "fmt"
func main() {
	num := -5
	switch {
	case num > 0:
		fmt.Println("Positive")
	case num < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
