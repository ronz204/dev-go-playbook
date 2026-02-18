package helpers

import "math/rand/v2"

func RandomInt(min, max int) int {
	return rand.IntN(max-min) + min
}
