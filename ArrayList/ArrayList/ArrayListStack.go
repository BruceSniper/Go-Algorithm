package ArrayList

type StackArray interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsEmpty() bool         //是否为空
	IsFull() bool          //是否满了
}

type Stack struct {
	myarray     *ArrayList
	capsize     int //最大范围
	currentsize int //实际使用大小
}

func NewArrayListStack() *Stack {
	mystack := new(Stack)
	mystack.myarray = NewArrayList() //数组
	mystack.capsize = 10
	mystack.currentsize = 0
	return mystack
}

func (mystack *Stack) Clear() {
	mystack.myarray.Clear()
}

func (mystack *Stack) Size() int {
	return mystack.myarray.TheSize
}

func (mystack *Stack) Pop() interface{} {
	if !mystack.IsEmpty() {
		last := mystack.myarray.dataStore[mystack.myarray.TheSize-1] //最后一个数据
		mystack.myarray.Delete(mystack.myarray.TheSize - 1)          //删除最后一个数据
		mystack.currentsize--
		return last
	}
	return nil
}

func (mystack *Stack) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data) //叠加数据
		mystack.currentsize++
	}
}

func (mystack *Stack) IsEmpty() bool {
	if mystack.myarray.TheSize == 0 {
		return true
	} else {
		return false
	}
}

func (mystack *Stack) IsFull() bool {
	if mystack.myarray.TheSize >= mystack.capsize {
		return true
	} else {
		return false
	}
}
