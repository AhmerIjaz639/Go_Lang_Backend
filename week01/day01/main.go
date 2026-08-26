package main

import "fmt"

const MyName = "Ahmer Ijaz"
const AppVersion = "1.0.0"

func main() {

    age := 25

    var city string = "Lahore"

    var isStudent bool = true

    var salary float64 = 500.5

    var zeroInt int
    var zeroString string
    var zeroBool bool

    fmt.Println("=== Personal Info ===")
    fmt.Printf("Name:       %s\n", MyName)
    fmt.Printf("Age:        %d\n", age)
    fmt.Printf("City:       %s\n", city)
    fmt.Printf("Student:    %t\n", isStudent)
    fmt.Printf("Salary:     %.2f\n", salary)

    fmt.Println("\n=== Types ===")
    fmt.Printf("MyName type:    %T\n", MyName)
    fmt.Printf("age type:       %T\n", age)
    fmt.Printf("city type:      %T\n", city)
    fmt.Printf("isStudent type: %T\n", isStudent)
    fmt.Printf("salary type:    %T\n", salary)

    fmt.Println("\n=== Zero Values ===")
    fmt.Printf("int zero value:    %d\n", zeroInt)
    fmt.Printf("string zero value: %q\n", zeroString)   // %q shows quotes around string
    fmt.Printf("bool zero value:   %t\n", zeroBool)
}