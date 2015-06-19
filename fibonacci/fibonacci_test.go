package fibonacci

import (
	"reflect"
	"testing"
)

var fibonacciSequence = []int{
	0,
	1,
	1,
	2,
	3,
	5,
	8,
	13,
	21,
	34,
	55,
	89,
	144,
	233,
	377,
	610,
	987,
	1597,
	2584,
	4181,
	6765,
	10946,
	17711,
	28657,
	46368,
	75025,
	121393,
	196418,
	317811,
	514229,
	832040,
	1346269,
	2178309,
	3524578,
	5702887,
	9227465,
	14930352,
	24157817,
	39088169,
}

func TestFibonacci(t *testing.T) {

	go func() {
		result := Fibonacci(-1)
		expected := 0
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Should be equal.\n\tExpected: %#v, Actual: %#v\n", expected, result)
		}
	}()

	go func() {
		result := Fibonacci(0)
		expected := 0
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Should be equal.\n\tExpected: %#v, Actual: %#v\n", expected, result)
		}
	}()

	go func() {
		result := Fibonacci(1)
		expected := 1
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Should be equal.\n\tExpected: %#v, Actual: %#v\n", expected, result)
		}
	}()

	go func() {
		for i := 2; i < len(fibonacciSequence); i++ {
			result := Fibonacci(i)
			expected := fibonacciSequence[i]
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Should be equal.\n\tExpected: %#v, Actual: %#v\n", expected, result)
			}
		}
	}()

}
