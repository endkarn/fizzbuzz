package fizzbuzz

import "strconv"

const (
	fizzNumber     = 3
	buzzNumber     = 5
	fizzBuzzNumber = 15
)

var fizzBuzzMap = make(map[int]string)

func init() {
	fizzBuzzMap[fizzBuzzNumber] = "FizzBuzz"
	fizzBuzzMap[buzzNumber] = "Buzz"
	fizzBuzzMap[fizzNumber] = "Fizz"
}

func FizzBuzz(number int) string {
	for keyMap, valueMap := range fizzBuzzMap {
		if number%keyMap == 0 {
			return valueMap
		}
	}
	return strconv.Itoa(number)
}
