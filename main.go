package main

import "leetcode/solutions"
import "fmt"

func main() {
	fmt.Println("Two sum")
	result := solutions.TwoSum([]int{1, 2, 3, 4, 4}, 8)
	fmt.Println(result)

	//fmt.Println("Reverse String")
	//str := "hello"
	//reverse := solutions.ReverseString([]byte(str))
	//fmt.Println(reverse)

	fmt.Println("Palindrome Number")
	isPalindrome := solutions.IsPalindrome(121)
	fmt.Println(isPalindrome)
	isPalindrome = solutions.IsPalindrome(-121)
	fmt.Println(isPalindrome)

	fmt.Println("Reverse Integer")
	revInt := solutions.Reverse(123)
	fmt.Println(revInt)
}
