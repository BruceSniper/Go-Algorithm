package main

import (
	"Go-Algorithm/StackArray/StackArray"
	"fmt"
)

func main() {
	mystack := StackArray.NewStack()
	mystack.Push(1)
	mystack.Push(2)
	mystack.Push(3)
	mystack.Push(4)
	//利用栈，逆序输出
	fmt.Println(mystack.Pop()) //4
	fmt.Println(mystack.Pop()) //3
	fmt.Println(mystack.Pop()) //2
	fmt.Println(mystack.Pop()) //1
}
