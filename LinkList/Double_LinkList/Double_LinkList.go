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
