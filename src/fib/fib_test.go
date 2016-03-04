package fib

import "testing"

func TestFibonacciZero(t *testing.T) {
	t.Log("Fibonacci(0) should return a sequence with only 0")
	sequence, _ := Fibonacci(0)
	if len(*sequence) != 1 {
		t.Errorf("Sequence len != 1. Result: %d", len(*sequence))
	} else if (*sequence)[0] != 0 {
		t.Errorf("sequence[0] should == 0. Result: %d", (*sequence)[0])
	}
}

func TestFibonacciOne(t *testing.T) {
	t.Log("Fibonacci(1) should return a sequence with 0 1")
	sequence, _ := Fibonacci(1)
	if len(*sequence) != 2 {
		t.Errorf("Sequence len != 2. Result: %d", len(*sequence))
	} else if (*sequence)[1] != 1 {
		t.Errorf("sequence[1] should == 1. Result: %d", (*sequence)[1])
	}
}

func TestFibonacciTwo(t *testing.T) {
	t.Log("Fibonacci(2) should return a sequence with 0 1 1")
	sequence, _ := Fibonacci(2)
	if len(*sequence) != 3 {
		t.Errorf("Sequence len != 3. Result: %d", len(*sequence))
	} else if (*sequence)[2] != 1 {
		t.Errorf("sequence[2] should == 2. Result: %d", (*sequence)[2])
	}
}

func TestFibonacciEight(t *testing.T) {
	t.Log("Fibonacci(8) should have 21 at element 8")
	sequence, _ := Fibonacci(8)
	if (*sequence)[8] != 21 {
		t.Errorf("sequence[8] should == 21. Result: %d", (*sequence)[8])
	}
}

func TestFibonacciNegative(t *testing.T) {
	t.Log("Error for unnatural number")
	_, err := Fibonacci(-1)
	if err == nil {
		t.Errorf("Expected to get an error.")
	}
}
