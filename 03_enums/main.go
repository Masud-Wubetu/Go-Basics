package main
import "fmt"
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	numbers := []int{10, 20, 30}
	fmt.Println(sum(numbers...))
}