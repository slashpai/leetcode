/*https://leetcode.com/problems/reverse-integer/*/

package solutions

func Reverse(x int) int {

	var neg = 1
	if x < 0 {
		neg = -1
		x = -1 * x
	}
	rev := 0
	for n := x; n > 0; n = n / 10 {
		rev = rev*10 + n%10

	}
	if rev < -(1<<31) || rev > ((1<<31)-1) {
		return 0
	}
	return rev * neg
}
