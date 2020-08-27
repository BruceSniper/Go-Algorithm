package Single_Link

import (
	"fmt"
)

//单链表的接口
type SingleLink interface {
	//增删查改
	GetFirstNode() *SingleLinkNode        //抓取头部节点
	InsertNodeFront(node *SingleLinkNode) //头部插入
	InsertNodeBack(node *SingleLinkNode)  //尾部插入

	InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool //根据dest值搜索，在该节点之前插入node
	InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool  //根据dest值搜索，在该节点之后插入node

	GetNodeAtIndex(index int) *SingleLinkNode //根据索引抓取指定位置的节点
	DeleteNode(dest *SingleLinkNode) bool     //删除一个节点，判断是否删除成功
	DeleteAtIndex(index int) bool             //根据索引值删除节点，判断是否成功
	String() string                           //返回链表字符串
	GetMid() *SingleLinkNode                  //求链表的中间位置
}

//单链表的结构
type SingleLinkList struct {
	head   *SingleLinkNode //链表头部指针
	length int             //链表长度
}

func NewSingleLinkList() *SingleLinkList {
	head := NewSingleLinkNode(nil)  //新建一个头结点
	return &SingleLinkList{head, 0} //返回一个链表
}

//返回以一个数据节点
func (list *SingleLinkList) GetFirstNode() *SingleLinkNode {
	return list.head.pNext
}

//头插法实现
func (list *SingleLinkList) InsertNodeFront(node *SingleLinkNode) {
	if list.head == nil {
		list.head = node
		node.pNext = nil
		list.length++ //插入节点，数据追加
	} else {
		p := list.head //备份头部节点位置
		node.pNext = p.pNext
		p.pNext = node
		list.length++ //插入节点，数据追加
	}

}

//头插法实现
func (list *SingleLinkList) InsertNodeBack(node *SingleLinkNode) {
	if list.head == nil {
		list.head = node
		node.pNext = nil
		list.length++ //插入节点，数据追加
	} else {
		p := list.head //备份头部位置信息
		for p.pNext != nil {
			p = p.pNext
		}
		p.pNext = node
		list.length++ //插入节点，数据追加
	}
}

func (list *SingleLinkList) String() string {
	var listString string
	p := list.head
	for p.pNext != nil {
		listString += fmt.Sprintf("%v--->", p.pNext.value)
		p = p.pNext
	}
	listString += fmt.Sprintf("nil")
	return listString //打印链表字符串
}

func (list *SingleLinkList) InsertNodeValueFront(dest interface{}, node *SingleLinkNode) bool {
	phead := list.head
	isfind := false
	for phead.pNext != nil {
		if phead.pNext.value == dest { //找到
			isfind = true
			break
		}
		phead = phead.pNext
	}
	if isfind {
		node.pNext = phead.pNext
		phead.pNext = node
		list.length++
		return true
	} else {
		return false
	}
}

func (list *SingleLinkList) InsertNodeValueBack(dest interface{}, node *SingleLinkNode) bool {
	phead := list.head
	isfind := false
	for phead != nil {
		if phead.value == dest { //找到
			isfind = true
			break
		}
		phead = phead.pNext
	}
	if isfind {
		node.pNext = phead.pNext
		phead.pNext = node
		list.length++
		return true
	} else {
		return false
	}

}

func (list *SingleLinkList) GetNodeAtIndex(index int) *SingleLinkNode {
	if index > list.length-1 || index < 0 {
		return nil
	} else {
		phead := list.head
		for index > -1 {
			phead = phead.pNext //向后循环
			index--             //向后循环的过程
		}
		return phead //返回的是地址
	}
}

func (list *SingleLinkList) DeleteNode(dest *SingleLinkNode) bool {
	if dest == nil {
		return false
	}
	phead := list.head
	for phead.pNext != nil && phead.pNext != dest {
		phead = phead.pNext //循环下去
	}
	if phead.pNext == dest {
		phead.pNext = phead.pNext.pNext
		list.length--
		return true
	} else {
		return false
	}
}

func (list *SingleLinkList) DeleteAtIndex(index int) bool {
	if index > list.length-1 || index < 0 {
		return false
	} else {
		phead := list.head
		for index > 0 {
			phead = phead.pNext //向后循环
			index--             //向后循环的过程
		}
		phead.pNext = phead.pNext.pNext
		list.length--
		return true
	}
}

func (list *SingleLinkList) GetMid() *SingleLinkNode {
	if list.head.pNext == nil {
		return nil
	} else {
		phead1 := list.head
		phead2 := list.head
		for phead2 != nil && phead2.pNext != nil {
			phead1 = phead1.pNext
			phead2 = phead2.pNext.pNext
		}
		return phead1 //中间节点
	}
}
