package fib

import "testing"

func TestFibonacciZero(t *testing.T) {
	t.Log("Fibonacci(0) should return a sequence with only 0")
	sequence := Fibonacci(0)
	if len(*sequence) != 1 {
		t.Errorf("Sequence len != 1. Result: %d", len(*sequence))
	} else if (*sequence)[0] != 0 {
		t.Errorf("sequence[0] should == 0. Result: %d", (*sequence)[0])
	}
}

func TestFibonacciOne(t *testing.T) {
	t.Log("Fibonacci(1) should return a sequence with 0 1")
	sequence := Fibonacci(1)
	if len(*sequence) != 2 {
		t.Errorf("Sequence len != 2. Result: %d", len(*sequence))
	} else if (*sequence)[1] != 1 {
		t.Errorf("sequence[1] should == 1. Result: %d", (*sequence)[1])
	}
}
