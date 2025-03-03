package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Пример использования:
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 5}

	// Вызов функции
	head = deleteDuplicates(head)

	// Выводим результат
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil{
        return head
    }
    dummy := &ListNode{Next: head}
    for head.Next != nil{
        if head.Next != nil && head.Val == head.Next.Val{
            if head.Next.Next != nil{
                head.Next = head.Next.Next
            }else{
                head.Next = nil
            }
        }
        if head.Next != nil && head.Val != head.Next.Val{
            head = head.Next
        }
    }
    return dummy.Next
}