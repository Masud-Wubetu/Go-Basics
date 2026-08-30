package main
import "fmt"
func main() {
	day := 4
	switch {
	case day == 1:
		fmt.Println("Monday")
	case day == 4:
		fmt.Println("Thursday")
	case day > 5:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}
}
