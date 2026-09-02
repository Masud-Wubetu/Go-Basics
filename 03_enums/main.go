package main
import (
    "fmt"
    "time"
)

func main() {

    switch hour := time.Now().Hour(); {
    case hour < 12:
        fmt.Println("Good Morning")
    case hour < 17:
        fmt.Println("Good Afternoon")
    default:
         fmt.Println("Good Evening")
    }

    checkType := func(i interface{}) {
        switch v := i.(type) {
        case int:
            fmt.Printf("Integer: %d\n", v)
        case string:
            fmt.Printf("String: %s\n", v)
        case bool:
            fmt.Printf("Boolean: %t\n", v)
        default:
            fmt.Printf("Unkown Type: %T\n", v) 
        }
    }

    checkType(21)
    checkType("Test")
    checkType(true)
    checkType(312.23)
    
}