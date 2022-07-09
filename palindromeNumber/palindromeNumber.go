package palindromeNumber

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	input := x

	if x < 0 {
		return false
	}

	rev := 0
	pop := 0

	for x != 0 {
		pop = x % 10
		x /= 10

		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			fmt.Println("Max int overflow")
		}

		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			fmt.Println("Min int overflow")
		}

		rev = rev*10 + pop
	}

	return input == rev
}
