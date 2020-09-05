/* https://leetcode.com/problems/reverse-string/*/

package solutions

func ReverseString(s []byte) {
	for i, j := 0, (len(s) - 1); i <= j; i, j = i+1, j-1 {
		t := s[i]
		s[i] = s[j]
		s[j] = t
	}
}
