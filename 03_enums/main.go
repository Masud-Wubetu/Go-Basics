package main
import "fmt"

func divide(a, b int) (int, int) {

    quotient := a / b
    remainder := a % b

    return quotient, remainder

}

func main() {
    q, r := divide(10, 3) 

    fmt.Println("quotient: ", q)
    fmt.Println("remainder: ", r)
}