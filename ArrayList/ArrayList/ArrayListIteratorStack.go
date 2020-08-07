package ArrayList

type StackArrayX interface {
	Clear()                //清空
	Size() int             //大小
	Pop() interface{}      //弹出
	Push(data interface{}) //压入
	IsEmpty() bool         //是否为空
	IsFull() bool          //是否满了
}

type StackX struct {
	myarray *ArrayList
	capsize int      //最大范围
	Myit    Iterator //迭代器
}

func NewArrayListStackX() *StackX {
	mystack := new(StackX)
	mystack.myarray = NewArrayList()          //数组
	mystack.Myit = mystack.myarray.Iterator() //迭代
	return mystack
}

func (mystack *StackX) Clear() {
	mystack.myarray.Clear()
	mystack.myarray.TheSize = 0
}

func (mystack *StackX) Size() int {
	return mystack.myarray.TheSize
}

func (mystack *StackX) Pop() interface{} {
	if !mystack.IsEmpty() {
		last := mystack.myarray.dataStore[mystack.myarray.TheSize-1] //最后一个数据
		mystack.myarray.Delete(mystack.myarray.TheSize - 1)          //删除最后一个数据
		return last
	}
	return nil
}

func (mystack *StackX) Push(data interface{}) {
	if !mystack.IsFull() {
		mystack.myarray.Append(data) //叠加数据
	}
}

func (mystack *StackX) IsEmpty() bool {
	if mystack.myarray.TheSize == 0 {
		return true
	} else {
		return false
	}
}

func (mystack *StackX) IsFull() bool {
	if mystack.myarray.TheSize >= 10 {
		return true
	} else {
		return false
	}
}
