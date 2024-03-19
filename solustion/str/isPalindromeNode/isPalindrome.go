package isPalindromeNode

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	str := ""
	for head != nil {
		str = strconv.Itoa(head.Val) + str
		head = head.Next
	}
	for i := 0; i < len(str)/2; i++ {
		fmt.Println(str[i], str[len(str)-1-i])
		if str[i] != str[len(str)-1] {
			return false
		}
	}
	return true
}
