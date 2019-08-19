package fizzbuzz

import "strconv"

const (
	FizzNumber     = 3
	BuzzNumber     = 5
	FizzBuzzNumber = 15
)

func FizzBuzz(number int) string {
	if number%FizzBuzzNumber == 0 {
		return "FizzBuzz"
	}
	if number%BuzzNumber == 0 {
		return "Buzz"
	}
	if number%FizzNumber == 0 {
		return "Fizz"
	}
	return strconv.Itoa(number)
}
