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

func Test_FizzBuzz_By_Number_4_Should_Be_4(t *testing.T) {
	expected := "4"
	number := 4

	actual := FizzBuzz(number)

	if expected != actual {
		t.Errorf("Expected %s but get %s .", expected, actual)
	}
}
