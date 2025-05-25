package pipeline

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	result := pipeline(numbers,
		triple,
		// addOne,
		filterEven,
	)

	fmt.Println(result) // Output: [4, 6, 8, 10]
}

// Pipeline step: transform
type transformFn func([]int) []int

func pipeline(data []int, steps ...transformFn) []int {
	for _, step := range steps {
		data = step(data)
	}
	return data
}

// --- Step 1: Double all numbers
func triple(numbers []int) []int {
	var res []int
	for _, n := range numbers {
		res = append(res, n*3)
	}
	return res
}

// --- Step 2: Add 1 to all numbers
func addOne(numbers []int) []int {
	var res []int
	for _, n := range numbers {
		res = append(res, n+1)
	}
	return res
}

// --- Step 3: Filter even numbers
func filterEven(numbers []int) []int {
	var res []int
	for _, n := range numbers {
		if n%2 == 0 {
			res = append(res, n)
		}
	}
	return res
}
