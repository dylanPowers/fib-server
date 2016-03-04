package fib

func Fibonacci(input int) *[]int {
	sequence := make([]int, 1)
	if input == 1 {
		sequence = append(sequence, 1)
	}

	return &sequence
}
