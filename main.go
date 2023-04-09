package main

import "fmt"

func cumulativeSum(numbers []int) []int {
	sum := 0
	cumSum := make([]int, len(numbers))

	for i, number := range numbers {
		sum += number
		cumSum[i] = sum
	}

	return cumSum
}

func main() {
	numbers := []int{1, 2, 3, 4}

	cumSum := cumulativeSum(numbers)

	fmt.Println(cumSum)
}