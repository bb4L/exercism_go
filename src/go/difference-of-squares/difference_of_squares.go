package diffsquares

// SquareOfSum returns the square of the sum
func SquareOfSum(n int) (result int) {
	result = (1 + n) * n / 2

	return result * result
}

// SumOfSquares return the sum of the squares
func SumOfSquares(n int) int {
	// https://www.tutorialspoint.com/sum-of-squares-of-first-n-natural-numbers-in-c-program
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference return difference of SqaresOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
