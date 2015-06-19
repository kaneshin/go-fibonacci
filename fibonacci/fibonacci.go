package fibonacci

const (
	n0 = 0
	n1 = 1
)

func Fibonacci(n int) int {
	if n < 1 {
		return n0
	}
	return f(n, n1, n0)
}

func f(n, a, b int) int {
	if n > 1 {
		return f(n-1, a+b, a)
	}
	return a
}
