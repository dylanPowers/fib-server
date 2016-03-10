package fib

import "errors"

// Fibonacci returns a splice containing the Fibonacci sequence for the given n.
// Any negative value will return an error.
func Fibonacci(n int) (*[]int, error) {
	if n < 0 {
		return nil, errors.New("Input is less than 0")
	}

	sequence := make([]int, 1)
	if n >= 1 {
		sequence = append(sequence, 1)
	}

	for i := 1; i < n; i++ {
		slen := len(sequence)
		sequence = append(sequence, sequence[slen-1]+sequence[slen-2])
	}

	return &sequence, nil
}
