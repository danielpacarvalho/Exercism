// Package diffsquares makes operations with integer number
package diffsquares

// SquareOfSum calculates the square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	resultado := 0
	for i := 0; i <= n; i++ {
		resultado += i
	}
	return resultado * resultado
}

// SquareOfSum calculates the sum of the squares  of the sum of the first N natural numbers
func SumOfSquares(n int) int {
	resultado := 0
	for i := 0; i <= n; i++ {
		resultado += (i * i)
	}
	return resultado
}

// Difference calculates the difference between the square of the sum of the first N natural numbers and the sum of the squares of said numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
