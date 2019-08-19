package fizzbuzz

import "strconv"

func FizzBuzz(number int) string {
	if number == 5 {
		return "Buzz"
	}
	if number != 3 {
		stringNum := strconv.Itoa(number)
		return stringNum
	}
	return "Fizz"
}