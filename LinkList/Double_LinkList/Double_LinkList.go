package Double_LinkList

import "fmt"

//双链表的基本结构
type DoubleLinkList struct {
	head   *DoubleLinkNode
	length int
}

//新建一个双链表
func NewDoubleLinkList() *DoubleLinkList {
	head := NewDoubleLinkNode(nil)
	return &DoubleLinkList{head, 0}
}

//返回链表长度
func (dlist *DoubleLinkList) GetLength() int {
	return dlist.length
}

//返回第一个节点
func (dlist *DoubleLinkList) GetFirstNode() *DoubleLinkNode {
	return dlist.head.next
}

//头部插入节点
func (dlist *DoubleLinkList) InsertHead(node *DoubleLinkNode) {
	phead := dlist.head
	if phead.next == nil {
		node.next = nil
		phead.next = node //只有一个节点直接连接上
		node.pre = phead  //上一个节点
		dlist.length++
	} else {
		phead.next.pre = node  //标记上一个节点
		node.next = phead.next //下一个节点
		phead.next = node
		node.pre = phead
		dlist.length++
	}
}

//尾部插入节点
func (dlist *DoubleLinkList) InsertBack(node *DoubleLinkNode) {
	phead := dlist.head
	if phead.next == nil {
		node.next = nil
		phead.next = node //只有一个节点直接连接上
		node.pre = phead  //上一个节点
		dlist.length++
	} else {
		for phead.next != nil {
			phead = phead.next //循环下去
		}
		phead.next = node //后缀
		node.pre = phead  //前缀
		dlist.length++
	}
}

func (dlist *DoubleLinkList) String() string {
	var listString1 string
	var listString2 string
	phead := dlist.head
	for phead.next != nil {
		//正向循环
		listString1 += fmt.Sprintf("%v--->", phead.next.value)
		phead = phead.next
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"

	for phead != dlist.head {
		//反向循环
		listString2 += fmt.Sprintf("<---%v", phead.value)
		phead = phead.pre
	}
	listString1 += fmt.Sprintf("nil")

	return listString1 + listString2 + "\n" //打印链表字符串

}

func (dlist *DoubleLinkList) InsertValueHead(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next != dest {
		phead = phead.next
	}
	if phead.next == dest {
		if phead.next.next != nil {
			phead.next.next.pre = node
		}
		node.next = phead.next
		node.pre = phead
		phead.next = node

		dlist.length++
		return true
	} else {
		return false
	}

}

func (dlist *DoubleLinkList) InsertValueBack(dest *DoubleLinkNode, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next != dest {
		phead = phead.next
	}
	if phead.next == dest {
		if phead.next.next != nil {
			phead.next.next.pre = node
		}
		node.next = phead.next.next
		phead.next.next = node
		node.pre = phead.next
		dlist.length++
		return true
	} else {
		return false
	}

}

func (dlist *DoubleLinkList) InsertValueHeadByValue(dest interface{}, node *DoubleLinkNode) bool {
	phead := dlist.head
	for phead.next != nil && phead.next.value != dest {
		phead = phead.next
	}
	if phead.next.value == dest {
		if phead.next.next != nil {
			phead.next.next.pre = node
		}
		node.next = phead.next
		node.pre = phead
		phead.next = node

		dlist.length++
		return true
	} else {
		return false
	}
}

func (dlist *DoubleLinkList) GetNodeAtIndex(index int) *DoubleLinkNode {
	if index > dlist.length-1 || index < 0 {
		return nil
	}
	phead := dlist.head
	for index > -1 {
		phead = phead.next
		index-- //计算位置
	}
	return phead
}

func (dlist *DoubleLinkList) DeleteNode(node *DoubleLinkNode) bool {
	if node == nil {
		return false
	} else {
		phead := dlist.head
		for phead.next != nil && phead.next != node {
			phead = phead.next
		}
		if phead.next == node {
			if phead.next.next != nil {
				phead.next.next.pre = phead //设置pre
			}
			phead.next = phead.next.next //设置next
			dlist.length--
			return true
		} else {
			return false
		}
	}
}

func (dlist *DoubleLinkList) DeleteNodeAtIndex(index int) bool {
	if index > dlist.length-1 || index < 0 {
		return false
	} else {
		phead := dlist.head
		for index > 0 {
			phead = phead.next
			index-- //计算位置
		}
		if phead.next.next != nil {
			phead.next.next.pre = phead //设置pre
		}
		phead.next = phead.next.next //设置next
		dlist.length--
		return true
	}
}
