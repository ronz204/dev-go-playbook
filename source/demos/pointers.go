package demos

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func PointersDemo() {
	number := 42
	direction := &number

	println("Value:", number)
	println("Address:", direction)
	println("Dereferenced value:", *direction)

	*direction = 100
	println("Updated value:", number)

	var alpha, beta int = 1, 2
	println("Before swap: alpha =", alpha, ", beta =", beta)

	Swap(&alpha, &beta)
	println("After swap: alpha =", alpha, ", beta =", beta)
}
