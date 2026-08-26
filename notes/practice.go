package main

import "fmt"

/*
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	println("Result:", result)
}

*/

func minMax(numbers []int) (min, max int) {
	min = numbers[0]
	max = numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}

	}

	return
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n

	}
	return total
}

func main() {
	/*numbers := []int{3, 5, 1, 8, 2}
	min, max := minMax(numbers)
	fmt.Printf("Min: %d, Max: %d\n", min, max)
	*/
    numbers := []int{3, 5, 1, 8, 2}

	fmt.Println("sum; ", sum(numbers...))

	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))

}
