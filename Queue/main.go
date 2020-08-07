package main

import (
	"Algorithm/Queue/Queue"
	"fmt"
)

func main() {
	myq := Queue.NewQueue()
	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue(3)
	myq.EnQueue(4)
	for i := 0; i < 6; i++ {
		fmt.Println(myq.DeQueue())
	}
}
