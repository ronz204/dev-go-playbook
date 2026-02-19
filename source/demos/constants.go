package demos

func ConstantsDemo() {
	const Pi float64 = 3.14159
	println("Value of Pi:", Pi)

	const (
		StatusOK       = 200
		StatusNotFound = 404
		StatusError    = 500
	)

	println("Status OK:", StatusOK)
	println("Status Not Found:", StatusNotFound)
	println("Status Error:", StatusError)
}
