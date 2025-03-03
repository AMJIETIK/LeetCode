package main

import "fmt"

func main() {
		fmt.Println(isPalindrome(9339))
}

func isPalindrome(x int) bool {
    check := x
	var palindrome int
	if x < 0 {
		return false
	}
	for check > 0 {
		palindrome = palindrome*10 + check%10
		check /= 10
	}
	if x == palindrome {
		return true
	}
	return false
}