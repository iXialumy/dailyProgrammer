package addToDigits

import "math"

func add2digits(input int) int {
	output := 0
	for i := 0; input > 0; i++ {
		lastdigit := input % 10
		output += int(math.Pow10(i)) * (lastdigit + 1)
		if lastdigit == 9 {
			i++
		}
		input = input / 10
	}

	return output
}
