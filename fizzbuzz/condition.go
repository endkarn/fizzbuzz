package fizzbuzz

import "strconv"

const (
	fizzNumber     = 3
	buzzNumber     = 5
	fizzBuzzNumber = 15
)

func FizzBuzz(number int) string {
	if number%fizzBuzzNumber == 0 {
		return "FizzBuzz"
	}
	if number%buzzNumber == 0 {
		return "Buzz"
	}
	if number%fizzNumber == 0 {
		return "Fizz"
	}
	return strconv.Itoa(number)
}
