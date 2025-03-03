//Definition for singly-linked list.
 package main

 import (
	"fmt"
 )

 type ListNode struct {
	Val int
	Next *ListNode
}

 func main() {
    // Создаем тестовый список сразу в main
    head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 3}
    head.Next.Next.Next = &ListNode{Val: 4}
    head.Next.Next.Next.Next = &ListNode{Val: 5}

    // Печатаем исходный список
    fmt.Println("Original list:")
    current := head
    for current != nil {
        fmt.Print(current.Val, " ")
        current = current.Next
    }
    fmt.Println()

    // Применяем reverseBetween
    result := reverseBetween(head, 2, 4)

    // Печатаем измененный список
    fmt.Println("Modified list:")
    current = result
    for current != nil {
        fmt.Print(current.Val, " ")
        current = current.Next
    }
    fmt.Println()
}

 func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || left == right {
        return head
    }

    reverseData := make([]int, 0, 5) 
    dummy := &ListNode{Next: head}
    prev := dummy

    for i := 1; i < left; i++ {
        prev = prev.Next
    }

    curr := prev.Next
    for i := left; i <= right; i++ {
        reverseData = append(reverseData, curr.Val)
        curr = curr.Next
    }

    for i := len(reverseData) - 1; i >= 0; i-- {
        prev.Next.Val = reverseData[i]
        prev = prev.Next
    }

    return dummy.Next
}
