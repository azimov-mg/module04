package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortAscending(slice []int) []int {
	sort.Ints(slice)
	return slice
}

func createList(nums []int) *ListNode {
	head := &ListNode{Val: nums[0], Next: nil}
	// сохраняем ссылку на текущий узел
	current := head
	// проходим по оставшимся элементам в nums, создавая новый узел для каждого из них
	for i := 1; i < len(nums); i++ {
		newNode := &ListNode{Val: nums[i], Next: nil}
		// устанавливаем ссылку на новый узел в качестве следующего узла текущего узла
		current.Next = newNode
		// переходим к следующему узлу
		current = newNode
	}
	// возвращаем головной узел (начало списка)
	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func printList(head *ListNode) {
	current := head
	
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

func main() {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	head := createList(sortAscending(s))

	deleteDuplicates(head)
	printList(head)
}
