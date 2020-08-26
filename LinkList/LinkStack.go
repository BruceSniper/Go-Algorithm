package LinkList

/*
	链式栈
*/

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
	if n.NextPoint == nil {
		return true
	} else {
		return false
	}
}

func (n *Node) Push(data interface{}) {
	newnode := &Node{Data: data}
	newnode.NextPoint = n.NextPoint
	n.NextPoint = newnode //头部插入
}

//栈操作方向，
func (n *Node) Pop() interface{} {
	if n.IsEmpty() == true {
		return nil
	}
	value := n.NextPoint.Data           //要弹出的数据
	n.NextPoint = n.NextPoint.NextPoint //删除
	return value
}

func (n *Node) Length() int {
	pnext := n
	length := 0
	for pnext.NextPoint != nil {
		pnext = pnext.NextPoint //节点的循环跳跃
		length++                //追加
	}
	return length //返回长度
}
