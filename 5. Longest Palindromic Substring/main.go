package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("bb"))
}

func longestPalindrome(s string) string {
	var answer string
	for i := 0; i<len(s); i++{
		odd := expandFromCenter(s, i, i)

		even := expandFromCenter(s, i, i+1)

		if len(odd) > len(answer){
			answer = odd
		}
		if len(even) > len(answer){
			answer = even
		}
	}
	return answer
}

func expandFromCenter(s string, left, right int) string{
	for left >= 0 && right < len(s) && s[left] == s[right]{
		left--
		right++
	}
	return s[left+1:right]
}
