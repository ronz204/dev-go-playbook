package demos

import (
	"errors"
	"playbook/source/helpers"
)

// sentinel errors
var SimpleError = errors.New("simple error")
var AnotherError = errors.New("another error")
var WrappedError = errors.New("wrapped error")

func launchError(number int) (string, error) {
	if number > 1 && number < 5 {
		return "", SimpleError

	} else if number > 5 && number < 10 {
		return "", AnotherError

	} else if number > 10 && number < 15 {
		return "", WrappedError

	} else {
		return "success", nil
	}
}

func ExceptionsDemo() {
	number := helpers.RandomInt(0, 20)
	result, err := launchError(number)

	if err != nil {
		switch err {
		case SimpleError:
			println("simple error occurred")
		case AnotherError:
			println("another error occurred")
		case WrappedError:
			println("wrapped error occurred")
		default:
			println("unknown error occurred")
		}
	} else {
		println("result:", result)
	}
}
