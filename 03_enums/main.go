package main

import "fmt"

func getNumbers() (int, int) {
	return 5, 10
}

func main() {

	a, b := getNumbers()

	fmt.Println("First Number: ", a)
	fmt.Println("Second Number: ", b)

}