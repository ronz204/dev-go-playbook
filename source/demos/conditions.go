package demos

import "fmt"

func ConditionsDemo() {
	if maybe := true; maybe {
		fmt.Println("Maybe is true")
	} else {
		fmt.Println("Maybe is false")
	}

	isWeekend, isHoliday, isRaining := true, false, true

	if isWeekend && !isHoliday && !isRaining {
		fmt.Println("It's a good day for a picnic!")
	}

	if isWeekend && (isHoliday || isRaining) {
		fmt.Println("It's a complicated day!")
	}

	if isWeekend || isHoliday {
		fmt.Println("It's a day off!")
	}

	if !isRaining {
		fmt.Println("It's not raining, you can go outside!")
	}
}
