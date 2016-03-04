package fib

func Fibonacci(input int) *[]int {
	sequence := make([]int, 1)
	if input >= 1 {
		sequence = append(sequence, 1)
	}

	for i := 1; i < input; i++ {
		slen := len(sequence)
		sequence = append(sequence, sequence[slen - 1] + sequence[slen - 2])
	}

	return &sequence
}
