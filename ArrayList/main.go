package main

import (
	"Algorithm/StackArray/StackArray"
	"fmt"
)

//func main() {
//	list := ArrayList.NewArrayList()
//	list.Append("a")
//	list.Append("b")
//	list.Append("c")
//	fmt.Println(list)
//	fmt.Println(list.TheSize) //小写是私有只能内部使用，大写是公有
//}

//func main() {
//	//定义接口对象，赋值的对象必须实现接口的所有方法
//	var list ArrayList.List = ArrayList.NewArrayList()
//	list.Append("a1")
//	list.Append("b2")
//	list.Append("c3")
//	fmt.Println(list)
//}

//func main() {
//	var list ArrayList.List = ArrayList.NewArrayList()
//	list.Append("a1")
//	list.Append("b2")
//	list.Append("c3")
//	list.Append("d3")
//	list.Append("f3")
//
//	for it := list.Iterator(); it.HasNext(); {
//		item, _ := it.Next()
//		if item == "d3" {
//			it.Remove()
//		}
//		fmt.Println(item)
//	}
//}
//func main() {
//	mystack := ArrayList.NewArrayListStackX()
//	//mystack.Push(1)
//	//mystack.Push(2)
//	//mystack.Push(3)
//	//mystack.Push(4)
//	////利用栈，逆序输出
//	//fmt.Println(mystack.Pop()) //4
//	//fmt.Println(mystack.Pop()) //3
//	//fmt.Println(mystack.Pop()) //2
//	//fmt.Println(mystack.Pop()) //1
//
//	for it := mystack.Myit;it.HasNext(); {
//		item, _ := it.Next()
//		fmt.Println(item)
//	}
//}

func main() {
	//模拟递归
	mystack := StackArray.NewStack()
	mystack.Push(5)
	last := 0 //保存结果
	for !mystack.IsEmpty() {
		data := mystack.Pop() //取出数据
		if data == 0 {
			last += 0
		} else {
			last += data.(int)
			mystack.Push(data.(int) - 1)
		}
	}
	fmt.Println(last)
}
