package LinkList

type Node struct {
	data  interface{}
	pNext *Node
}

type LinkList struct {
	head   *Node
	length uint
}

func NewList() *Node {
	return &Node{} //返回一个节点指针
}

func Find(n *Node, data interface{}) *Node {
	tmp := n
	if tmp.data != data {
		tmp = tmp.pNext
	}
	return tmp
}

// 链表是否为空
func IsEmpty(n *Node) bool {
	return n.pNext == nil
}

// 是否是最后一个节点
func IsLast(n *Node) bool {
	return n.pNext == nil
}

// 插入节点:在指定节点插入节点
/**
position:指定的节点位置
*/
func Insert(data interface{}, position *Node) {
	//新建一个节点
	tmpNode := new(Node)
	// 给新建的节点的值域赋值为传入的data值
	tmpNode.data = data
	// 新建的节点的next指针指向的是指定节点position的next
	tmpNode.pNext = position.pNext
	// 指定节点position的后继节点变成了新建的节点
	position.pNext = tmpNode
}

// 查找指定节点的前一个节点
func FindPrevious(data interface{}, n *Node) *Node {
	tmp := n
	for tmp.pNext != nil && tmp.pNext.data != data {
		tmp = tmp.pNext
	}
	return tmp
}

// 删除节点
func Delete(data interface{}, n *Node) {

}

//func (n *Node) Push(data interface{}) {
//	newnode := &Node{data: data}
//	newnode.pNext = n.pNext
//	n.pNext = newnode //头部插入
//}
//
////栈操作方向，
//func (n *Node) Pop() interface{} {
//	if n.IsEmpty() == true {
//		return nil
//	}
//	value := n.pNext.data   //要弹出的数据
//	n.pNext = n.pNext.pNext //删除
//	return value
//}
//
//func (n *Node) Length() int {
//	pnext := n
//	length := 0
//	for pnext.pNext != nil {
//		pnext = pnext.pNext //节点的循环跳跃
//		length++            //追加
//	}
//	return length //返回长度
//}
