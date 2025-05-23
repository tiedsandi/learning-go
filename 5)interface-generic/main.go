package main

import (
	"fmt"
)

type Addable interface {
	~int | ~float64 | ~string
}

func main() {
	result := add2(1, 2)

	fmt.Println(result)

	rawResult, err := add11(3, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		result := rawResult.(int) // type assertion: yakinkan bahwa dia int
		fmt.Println(result)
	}

	res, err := add11(3, 4)
	if err != nil {
		fmt.Println(err)
		return
	}

	intRes, ok := res.(int)
	if !ok {
		fmt.Println("Result is not an int!")
		return
	}

	fmt.Println("Result:", intRes)

}

// func add1[T interface{}](a, b T) T {
// 	return a + b
// }

func add11[T any](a, b T) (any, error) {
	switch aTyped := any(a).(type) {
	case int:
		bTyped, ok := any(b).(int)
		if ok {
			return aTyped + bTyped, nil
		}
	case float64:
		bTyped, ok := any(b).(float64)
		if ok {
			return aTyped + bTyped, nil
		}
	case string:
		bTyped, ok := any(b).(string)
		if ok {
			return aTyped + bTyped, nil
		}
	}

	return nil, fmt.Errorf("invalid type for + operation: %T", a)
}

func add2[T int | float64 | string](a, b T) T {
	return a + b
}

// func addAndChange[T any, K any](a, b T) K {
// 	aInt, aIsInt := a.(int)
// 	bInt, bIsInt := b.(int)

// 	if aIsInt && bIsInt {
// 		return aInt + bInt
// 	}

// 	aFloat, aIsFloat := a.(float64)
// 	bFloat, bIsFloat := b.(float64)

// 	if aIsFloat && bIsFloat {
// 		return aFloat + bFloat
// 	}

// 	return nil
// }

func add3[T Addable](a, b T) T {
	return a + b
}
