package checkDigitsUPC

import "math"

func checkDigitsUPC(input int) int {
	if input > 99999999999 {
		return -1
	}

	checkDigit := 0
	// 1. Sum up the digits at even numbered positions
	for i := 0; i <= 10; i = i + 2 {
		checkDigit += (input / int(math.Pow10(i))) % 10
	}

	// 2. Multiply by 3
	checkDigit *= 3

	// 3. Take the sum of digits at even-numbered positions in the original number, and add this sum to the result from step 2.
	for i := 1; i <= 11; i = i + 2 {
		checkDigit += (input / int(math.Pow10(i))) % 10
	}

	// 4. Find the result from step 3 modulo 10 (i.e. the remainder, when divided by 10) and call it M.
	m := checkDigit % 10

	// 5. If M is 0, then the check digit is 0; otherwise the check digit is 10 - M.
	if m == 0 {
		return 0
	}

	return 10 - m
}
