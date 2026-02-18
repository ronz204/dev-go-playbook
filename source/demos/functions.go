package demos

import "fmt"

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, nil
}

func reduce(numbers ...int) (total int) {
	for _, number := range numbers {
		total += number
	}
	return
}

func FunctionsDemo() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 divided by 2 is", result)
	}

	sum := reduce(1, 2, 3, 4, 5)
	fmt.Println("Sum of 1,2,3,4,5 is", sum)
}
