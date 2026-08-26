package main

import (
	"errors"
	"fmt"
)

func greet(name string) string {
	return "hello, " + name
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not possible")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {

		fmt.Println("Error:", err)
		return
	}
	fmt.Println("result:", result)

}
