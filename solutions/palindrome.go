/*https://leetcode.com/problems/palindrome-number/*/

package solutions

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		rev := 0
		for n := x; n > 0; n = n / 10 {
			rev = rev*10 + n%10

		}
		return rev == x
	}

}
