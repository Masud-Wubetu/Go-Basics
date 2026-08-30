package main
import "fmt"
func main() {
	var value interface{} = 3.14
	switch v := value.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float:", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
