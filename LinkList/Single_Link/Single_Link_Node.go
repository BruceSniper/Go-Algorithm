package Single_Link

//单链表的节点
type SingleLinkNode struct {
	value interface{}
	pNext *SingleLinkNode
}

//构造一个节点
func NewSingleLinkNode(data interface{}) *SingleLinkNode {
	return &SingleLinkNode{data, nil}
}

//返回数据
func (node *SingleLinkNode) Value() interface{} {
	return node.value
}

//返回下一个节点
func (node *SingleLinkNode) PNext() *SingleLinkNode {
	return node.pNext
}
