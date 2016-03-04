package fib

import "errors"

func Fibonacci(input int) (*[]int, error) {
	if input < 0 {
		return nil, errors.New("Input is less than 0")
	}

	sequence := make([]int, 1)
	if input >= 1 {
		sequence = append(sequence, 1)
	}

	for i := 1; i < input; i++ {
		slen := len(sequence)
		sequence = append(sequence, sequence[slen-1]+sequence[slen-2])
	}

	return &sequence, nil
}
