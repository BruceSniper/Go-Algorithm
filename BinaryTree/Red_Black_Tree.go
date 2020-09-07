package main

type (
	color uint //颜色

	//添加的元素
	Element interface {
		Value() interface{}
		Compare(Element) int //相等返回0，小于返回负数，大于返回正数
	}

	//红黑树节点
	RedBlackNode struct {
		Color  color         //节点颜色
		Data   Element       //数据
		Parent *RedBlackNode //父亲节点
		Left   *RedBlackNode //左孩子
		Right  *RedBlackNode //右孩子
	}

	RedBlackTree struct {
		Root *RedBlackNode //根节点
	}
)

const (
	Black color = iota + 1
	Red
)

//获取某个节点的祖父节点（爷爷节点）
func (t *RedBlackNode) grandParent() *RedBlackNode {
	//根节点没有父节点和祖父节点
	if t.Parent == nil {
		return nil
	}
	return t.Parent.Parent
}

//获取某个节点的叔父节点（父亲节点的兄弟节点）
func (t *RedBlackNode) uncleNode() *RedBlackNode {
	if t.grandParent() == nil {
		return nil
	}
	if t.Parent == t.grandParent().Left {
		return t.grandParent().Right
	}
	return t.grandParent().Left
}

func insertCase1(node *RedBlackNode) (root *RedBlackNode) {
	if node.Parent == nil {
		node.Parent = Black
		root = node
		return
	}
	return insertCase2(node)
}

func insertCase2(node *RedBlackNode) (root *RedBlackNode) {
	if node.Parent.Color == Black {
		return
	}
	return insertCase3(node)
}

func insertCase3(node *RedBlackNode) (root *RedBlackNode) {
	if uncle := node.uncleNode(); uncle != nil && uncle.Color == Red {
		grandParent := node.grandParent()
		node.Parent.Color = Black
		uncle.Color = Black
		grandParent.Color = Red
		return insertCase1(grandParent)
	}
	return insertCase4(node)
}

func insertCase4(node *RedBlackNode) (root *RedBlackNode) {
	grandParent := node.grandParent()
	if node == node.Parent.Right && node.Parent == grandParent.Left {
		leftRotate(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandParent.Right {
		rightRotate(node.Parent)
		node = node.Right
	}
	return insertCase5(node)
}

func insertCase5(node *RedBlackNode) (root *RedBlackNode) {
	grandParent := node.grandParent()
	node.Parent.Color = Black
	grandParent.Color = Red
	if node == node.Parent.Left && node.Parent == grandParent.Left {
		root = rightRotate(grandParent)
	} else {
		root = leftRotate(grandParent)
	}
	return
}

/*
红黑树删除节点
*/

//找到某个节点的兄弟节点
func (t *RedBlackNode) siblingNode() *RedBlackNode {
	if t.Parent == nil {
		return nil
	}
	if t == t.Parent.Left {
		return t.Parent.Right
	}
	return t.Parent.Left
}

//删除node节点，这里被删除的节点一定最多只有一个非叶子节点的孩子节点
func deleteNode(node *RedBlackNode) (root *RedBlackNode) {
	var child *RedBlackNode
	if node.Left.Data == nil {
		child = node.Right
	} else {
		child = node.Left
	}
	//如果删除的是没有子树的节点
	if node.Parent == nil && node.Left.Data == nil && node.Right.Data == nil {
		node.Data = nil
		node.Color = Black
		root = node
		return
	}
	//将孩子节点提升，这里已经将node删除
	if node.Parent.Left == node {
		node.Parent.Left = child
	} else {
		node.Parent.Right = child
	}
	child.Parent = node.Parent

	//如果node为红色，删除红色节点不影响红黑树的性质，不需要调整
	//如果node为黑色，但孩子节点为红色，此时只需要将孩子节点改为黑色即可
	//如果node为黑色，孩子节点也是黑色，则少了一个黑色节点，需要对树进行调整，一达到平衡
	if node.Color == Black {
		if child.Color == Red { //如果删除的是黑色节点并且提升的孩子节点为红色，则只需要将孩子节点调色为黑色即可
			child.Color = Black
		} else { //否则需要其他操作
			root = deleteCase1(child)
		}
	}
	return
}

func deleteCase1(node *RedBlackNode) (root *RedBlackNode) {
	if node.Parent == nil {
		return node
	}
	return deleteCase2(node)
}

func deleteCase2(node *RedBlackNode) (root *RedBlackNode) {
	var siblingNode = node.siblingNode()
	if siblingNode == nil {
		return
	}
	if siblingNode.Color == Red {
		node.Parent.Color = Red
		siblingNode.Color = Black
		if node == node.Parent.Left {
			leftRotate(node.Parent)
		} else {
			rightRotate(node.Parent)
		}
	}
	return deleteCase3(node)
}

func deleteCase3(node *RedBlackNode) (root *RedBlackNode) {
	var siblingNode = node.siblingNode()
	if (node.Parent.Color == Black) &&
		(siblingNode.Color == Black) &&
		(siblingNode.Left.Color == Black) &&
		(siblingNode.Right.Color == Black) {
		siblingNode.Color = Red
		return deleteCase1(node.Parent)
	} else {
		return deleteCase4(node)
	}
}

func deleteCase4(node *RedBlackNode) (root *RedBlackNode) {
	var siblingNode = node.siblingNode()
	if (node.Parent.Color == Red) &&
		(siblingNode.Color == Black) &&
		(siblingNode.Left.Color == Black) &&
		(siblingNode.Right.Color == Black) {
		siblingNode.Color = Red
		node.Parent.Color = Black
	} else {
		root = deleteCase5(node)
	}
	return
}

func deleteCase5(node *RedBlackNode) (root *RedBlackNode) {
	var siblingNode = node.siblingNode()

	if siblingNode.Color == Black {
		if (node == node.Parent.Left) &&
			(siblingNode.Right.Color == Black) &&
			(siblingNode.Left.Color == Red) {
			siblingNode.Color = Red
			siblingNode.Right.Color = Black
			rightRoate(siblingNode)
		} else if (node == node.Parent.Right) &&
			(siblingNode.Left.Color == Black) &&
			(siblingNode.Right.Color == Red) {
			siblingNode.Color = Red
			siblingNode.Right.Color = Black
			leftRotate(siblingNode)
		}
	}
	return deleteCase6(node)
}

func deleteCase6(node *RedBlackNode) (root *RedBlackNode) {
	var siblingNode = node.siblingNode()
	siblingNode.Color = node.Parent.Color
	node.Parent.Color = Black

	if node == node.Parent.Left {
		siblingNode.Right.Color = Black
		root = leftRotate(node.Parent)
	} else {
		siblingNode.Left.Color = Black
		root = rightRotate(node.Parent)
	}
	return
}
