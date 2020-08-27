package Single_Link

import "fmt"

//单链表的接口
type SingleLink interface {
	//增删查改
	GetFirstNode() *SingleLinkNode            //抓取头部节点
	InsertNodeFront(node *SingleLinkNode)     //头部插入
	InsertNodeBack(node *SingleLinkNode)      //尾部插入
	GetNodeAtIndex(index int) *SingleLinkNode //根据索引抓取指定位置的节点
	DeleteNode(dest *SingleLinkNode) bool     //删除一个节点，判断是否删除成功
	DeleteAtIndex(index int) bool             //根据索引值删除节点，判断是否成功
	String() string                           //返回链表字符串
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
