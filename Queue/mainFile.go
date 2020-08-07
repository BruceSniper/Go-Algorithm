package main

import (
	"Algorithm/Queue/Queue"
	"fmt"
	"io/ioutil"
)

//用队列来实现文件遍历
func main() {
	path := "D:\\codes" //路径
	var files []string  //数组字符串

	myqueue := Queue.NewQueue()
	myqueue.EnQueue(path)

	for {
		path := myqueue.DeQueue() //不断从队列中取出数据
		if path == nil {
			break
		}
		fmt.Println("get ", path)
		read, _ := ioutil.ReadDir(path.(string)) //读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path.(string) + "\\" + fi.Name() //构造新的路径
				fmt.Println("Dir ", fulldir)                //文件夹
				myqueue.EnQueue(fulldir)
			} else {
				fulldir := path.(string) + "\\" + fi.Name() //构造新的路径
				files = append(files, fulldir)              //追加路径
				fmt.Println("File ", fulldir)               //文件
			}
		}

	}
	for i := 0; i < len(files); i++ { //打印
		fmt.Println(files[i])
	}
}
