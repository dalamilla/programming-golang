// 2520 is the smallest number that can be divided by each of
// the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible
// by all of the numbers from 1 to 20?

package euler

// Solution of fifth Euler problem.
func Euler005(n int) int {
	found := true
	number := 0

	for found {
		i := 1
		number += n

		for number%i == 0 && i <= n {
			if i == n {
				found = false
			}
			i += 1
		}
	}

	return number
}
