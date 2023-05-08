package fizzbuzz

import "strconv"

func FizzBuzz(token string) string {
	integer, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		return "Error"
	}
	switch {
	case integer%15 == 0:
		return "FizzBuzz"
	case integer%3 == 0:
		return "Fizz"
	case integer%5 == 0:
		return "Buzz"
	default:
		return token
	}
}
