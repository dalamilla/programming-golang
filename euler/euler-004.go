// A palindromic number reads the same both ways. The largest palindrome made from 
// the product of two 2-startDigit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-startDigit numbers.

package euler

import "math"

// Solution of fourth Euler problem.
func Euler004(n int) int {
	startDigit := int(math.Pow(10, float64(n-1)))
	endDigit := int(math.Pow(10, float64(n)))
	max := 0

	for i := startDigit; i < endDigit; i++ {
		for j := startDigit; j < endDigit; j++ {
			if isPalindrome(i*j) &&  i * j > max {
					max = i * j
			}
		}
	}

	return max
}

// Helper function of fourth Euler problem.
func isPalindrome(n int) bool {
	rev_num := 0
	ori_num := n
	
	for n > 0 {
			remainder := n % 10
			rev_num *= 10
			rev_num += remainder 
			n /= 10
	}

	return ori_num == rev_num 
}
