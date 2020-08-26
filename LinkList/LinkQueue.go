package LinkList

/*
	链式队列
*/

type QueueLink struct {
	rear  *Node
	front *Node
}

type LinkQueue interface {
	length() int
	Enqueue(value interface{})
	Dequeue() interface{}
}

func NewLinkQueue() *QueueLink {
	return &QueueLink{}
}

func (qlk *QueueLink) length() int {
	pnext := qlk.front
	length := 0
	for pnext.NextPoint != nil {
		pnext = pnext.NextPoint //节点循环
		length++                //追加
	}
	return length //返回长度
}

func (qlk *QueueLink) Enqueue(value interface{}) {
	newnode := &Node{value, nil, nil} //新节点
	if qlk.front == nil {
		qlk.front = newnode //插入一个节点
		qlk.rear = newnode
	} else {
		qlk.rear.NextPoint = newnode
		qlk.rear = qlk.rear.NextPoint
	}
}

func (qlk *QueueLink) Dequeue() interface{} {
	if qlk.front == nil {
		return nil
	}
	newnode := qlk.front       //记录头部位置
	if qlk.front == qlk.rear { //只剩下一个
		qlk.front = nil
		qlk.rear = nil
	} else {
		qlk.front = qlk.front.NextPoint //删除一个
	}
	return newnode.Data
}
