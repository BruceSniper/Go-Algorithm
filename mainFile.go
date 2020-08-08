package main

import (
	"Go-Algorithm/StackArray/StackArray"
	"errors"
	"fmt"
	"io/ioutil"
)

/*
	循环获取所有文件夹的目录、内容
*/
func main1() {
	path := "D:\\codes"               //路径
	var files []string                //数组字符串
	files, _ = GetAll(path, files)    //抓取所有文件
	for i := 0; i < len(files); i++ { //打印路径
		fmt.Println(files[i])
	}
}

func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, fi := range read { //循环每个文件或文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "\\" + fi.Name() //构造新的路径
			files = append(files, fulldir)     //追加路径
			files, _ = GetAll(fulldir, files)  //文件夹递归处理
		} else {
			fulldir := path + "\\" + fi.Name() //构造新的路径
			files = append(files, fulldir)     //追加路径
		}
	}
	return files, nil
}

func main() {
	path := "D:\\codes"              //路径
	var files []string               //数组字符串
	mystack := StackArray.NewStack() //初始化一个栈
	mystack.Push(path)               //压入栈
	for !mystack.IsEmpty() {
		path := mystack.Pop().(string)
		files = append(files, path)     //加入列表
		read, _ := ioutil.ReadDir(path) //读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path + "\\" + fi.Name() //构造新的路径
				//files = append(files, fulldir)     //追加路径
				mystack.Push(fulldir) //文件夹递归处理
			} else {
				fulldir := path + "\\" + fi.Name() //构造新的路径
				files = append(files, fulldir)     //追加路径
			}
		}
	}
	for i := 0; i < len(files); i++ { //打印文件夹里的文件
		fmt.Println(files[i])
	}
}
