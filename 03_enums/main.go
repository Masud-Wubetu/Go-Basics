package main

import "fmt"
import "rsc.io/quote"

const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday

)

func main() {

	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Thursday)
	fmt.Println(Wednesday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(quote.Go())

}