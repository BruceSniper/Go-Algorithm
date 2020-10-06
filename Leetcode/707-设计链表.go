package Leetcode

type MyLinkedList struct {
	Header *ListNode `json:"header"`
	Tail   *ListNode `json:"tail"`
	Lens   int       `json:"lens"`
}

/** Initialize your data structure here. */
func LinkedListConstructor() MyLinkedList {
	return MyLinkedList{
		Header: nil,
		Tail:   nil,
		Lens:   0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (ll *MyLinkedList) Get(index int) int {
	// 如果获取的位置小于0或者等于链表长度,直接返回-1(注意链表下标从0开始,所以这地方可以等于)
	if index < 0 || index >= ll.Lens {
		return -1
	}

	// 如果index等于0,直接返回头节点的值
	if index == 0 {
		return ll.Header.Val
	}

	// 遍历一下,找到index节点的值
	node := ll.Header
	for node.Next != nil {
		// 因为0的情况一排除,所以直接先减掉
		index--
		// node指针往下移动一位
		if node.Next != nil {
			node = node.Next
		}
		// 当index递减等于0的时候, 返回其值就可以了
		if index == 0 {
			return node.Val
		}
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (ll *MyLinkedList) AddAtHead(val int) {
	ll.Header = &ListNode{
		Val:  val,
		Next: ll.Header,
	}
	// 如果当前链表为空,那么增加一个节点,这个节点既是头节点又是尾节点
	if ll.Lens == 0 {
		ll.Tail = ll.Header
	}
	// 因为增加了节点,所以链表长度+1
	ll.Lens++

}

/** Append a node of value val to the last element of the linked list. */
func (ll *MyLinkedList) AddAtTail(val int) {
	// 如果当前链表为空,那么增加尾部,也就是加个头部..
	if ll.Lens == 0 {
		ll.Tail = &ListNode{
			Val:  val,
			Next: nil,
		}
		ll.Header = ll.Tail
		ll.Lens++
		return
	}
	// 尾节点本来next等于nil,现在加一个,next等于这个节点
	ll.Tail.Next = &ListNode{
		Val:  val,
		Next: nil,
	}

	// 所以以后新的尾节点就是之前的next节点了..
	ll.Tail = ll.Tail.Next

	// 新增节点,链表长度+1
	ll.Lens++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (ll *MyLinkedList) AddAtIndex(index int, val int) {
	//   如果 index小于0，则在头部插入节点。
	if index <= 0 {
		ll.AddAtHead(val)
		return
	}

	//   如果 index 大于链表长度，则不会插入节点。
	if index > ll.Lens {
		return
	}

	//   如果 index 等于链表的长度，则该节点将附加到链表的末尾。
	if index == ll.Lens {
		ll.AddAtTail(val)
		return
	}

	node := ll.Header
	for node.Next != nil {
		index--
		// 当index == 0的时候,说明找到了这个节点,往这节点之前插入节点
		if index == 0 {
			newNode := &ListNode{
				Val:  val,
				Next: node.Next,
			}
			node.Next = newNode
			// 记得长度+1
			ll.Lens++
			// 记得要返回..
			return
		}

		node = node.Next
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (ll *MyLinkedList) DeleteAtIndex(index int) {
	// 如果index小于0或者大于等于长度,直接返回
	if index < 0 || index >= ll.Lens {
		return
	}

	// 如果等于0,就是删除头节点,记得链表长度-1
	if index == 0 {
		ll.Header = ll.Header.Next
		ll.Lens--
	}

	node := ll.Header
	for node.Next != nil {
		index--

		if index == 0 {
			// 如果node.Next.Next == nil 说明到最后一个节点了.相当于删除最后一个节点
			if node.Next.Next == nil {
				node.Next = nil
				ll.Tail = node
				ll.Lens--
				return
			}
			// 其他情况就是删除中间一个节点(A->B->C),操作就是  A 直接指向 C 就行 (A->C)
			node2 := node.Next.Next
			node.Next = node2
			ll.Lens--
			return
		}
		node = node.Next
	}
}
