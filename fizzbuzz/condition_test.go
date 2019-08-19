package fizzbuzz

import "testing"

func Test_FizzBuzz_By_Number_3_Should_Be_Fizz(t *testing.T) {
	expected := "Fizz"
	number := 3

	actual := FizzBuzz(number)

	if expected != actual {
		t.Errorf("Expected %s but get %s .", expected, actual)
	}
}
