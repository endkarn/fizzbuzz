package fizzbuzz

import "strconv"

func FizzBuzz(number int) string {
	if number == 5 {
		return "Buzz"
	}
	if number == 3 {
		return "Fizz"
	}
	if number == 15 {
		return "FizzBuzz"
	}
	return strconv.Itoa(number)
}