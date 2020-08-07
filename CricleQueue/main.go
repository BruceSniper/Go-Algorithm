package main

import (
	"Algorithm/CricleQueue/CricleQueue"
	"fmt"
)

func main() {
	var myq CricleQueue.CricleQueue

	CricleQueue.InitQueue(&myq)
	CricleQueue.EnQueue(&myq, 1)
	CricleQueue.EnQueue(&myq, 2)
	CricleQueue.EnQueue(&myq, 3)
	CricleQueue.EnQueue(&myq, 4)
	CricleQueue.EnQueue(&myq, 5)

	fmt.Println(CricleQueue.DeQueue(&myq))
	fmt.Println(CricleQueue.DeQueue(&myq))
	fmt.Println(CricleQueue.DeQueue(&myq))
	fmt.Println(CricleQueue.DeQueue(&myq))
	fmt.Println(CricleQueue.DeQueue(&myq))

}
