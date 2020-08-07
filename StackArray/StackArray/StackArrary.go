package StackArray

type StackArray interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsEmpty() bool         //是否为空
	IsFull() bool          //是否满了
}

type Stack struct {
	dataSource  []interface{}
	capsize     int //最大范围
	currentsize int //实际使用大小
}

func NewStack() *Stack {
	mystack := new(Stack)
	mystack.dataSource = make([]interface{}, 0, 1000)
	mystack.capsize = 1000
	mystack.currentsize = 0
	return mystack
}

func (mystack *Stack) Clear() {
	mystack.dataSource = make([]interface{}, 0, 1000) //重新开辟内存
	mystack.currentsize = 0                           //大小重设为0
	mystack.capsize = 1000
}

func (mystack *Stack) Size() int {
	return mystack.currentsize
}

func (mystack *Stack) Pop() interface{} {
	if !mystack.IsEmpty() {
		last := mystack.dataSource[mystack.currentsize-1]               //最后一个数据
		mystack.dataSource = mystack.dataSource[:mystack.currentsize-1] //删除最后一个数据
		mystack.currentsize--
		return last
	}
	return nil
}

func (mystack *Stack) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.dataSource = append(mystack.dataSource, data) //叠加数据
		mystack.currentsize++
	}
}

func (mystack *Stack) IsEmpty() bool {
	if mystack.currentsize == 0 {
		return true
	} else {
		return false
	}
}

func (mystack *Stack) IsFull() bool {
	if mystack.currentsize >= mystack.capsize {
		return true
	} else {
		return false
	}
}
