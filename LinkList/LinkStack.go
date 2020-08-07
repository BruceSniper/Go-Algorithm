package main

/*
	链式栈
*/
type Node struct {
	data  interface{}
	pNext *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewStack() *Node {
	return &Node{} //返回一个节点指针
}

func (n *Node) IsEmpty() bool {
	if n.pNext == nil {
		return true
	} else {
		return false
	}

}

func (n *Node) Push(data interface{}) {
	newnode := &Node{data: data}
	newnode.pNext = n.pNext
	n.pNext = newnode //头部插入
}

//栈操作方向，
func (n *Node) Pop() interface{} {
	if n.IsEmpty() == true {
		return nil
	}
	value := n.pNext.data   //要弹出的数据
	n.pNext = n.pNext.pNext //删除
	return value
}

func (n *Node) Length() int {
	pnext := n
	length := 0
	for pnext.pNext != nil {
		pnext = pnext.pNext //节点的循环跳跃
		length++            //追加
	}
	return length //返回长度
}
