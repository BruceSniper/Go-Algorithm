package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	List := make([]*ListNode, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &List[i].Val)
		List[i] = List[i].Next
	}
	result := mergeKLists(List)
	for i := 0; result.Next != nil; i++ {
		fmt.Printf("%d ", &result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	guard := &ListNode{}
	for i := 0; i < len(lists); i++ {
		guard.Next = mergeTwoLists(guard.Next, lists[i])
	}
	return guard.Next
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	guard := &ListNode{}
	head := guard
	for {
		if l1 == nil {
			head.Next = l2
			break
		}
		if l2 == nil {
			head.Next = l1
			break
		}

		if l1.Val < l2.Val {
			v := l1.Val
			for l1 != nil && v == l1.Val {
				node := l1
				l1 = l1.Next
				head.Next = node
				head = head.Next
			}
		} else {
			v := l2.Val
			for l2 != nil && v == l2.Val {
				node := l2
				l2 = l2.Next
				head.Next = node
				head = head.Next
			}
		}
	}
	return guard.Next
}
