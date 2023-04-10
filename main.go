//Write a program that takes a list of integers as input and returns a new list containing the cumulative sum of the original list (i.e., the running total of the numbers). For example, given the input list [1, 2, 3, 4], the output should be [1, 3, 6, 10].


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
