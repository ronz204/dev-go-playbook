package demos

import "fmt"

func CompositeDemo() {
	// Arrays
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	numbers[0] = 10
	fmt.Println("Updated Numbers:", numbers)

	// Slices
	var letters = make([]rune, 5, 11)
	letters[0] = 'H'
	letters[1] = 'e'
	letters[2] = 'l'
	letters[3] = 'l'
	letters[4] = 'o'
	fmt.Println("Letters:", string(letters))

	// Maps
	var people = make(map[string]int)
	people["Alice"] = 30
	people["Bob"] = 25
	fmt.Println("People:", people)

	var products = map[string]float64{
		"Laptop": 999.99,
		"Phone":  499.49,
		"Tablet": 299.29,
	}

	// Comma ok idiom to check if a key exists in the map
	if price, exists := products["Laptop"]; exists {
		fmt.Println("Laptop Price:", price)
	}
}
