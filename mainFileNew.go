package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func main() {
	path := "D:\\codes"                //路径
	var files []string                 //数组字符串
	files, _ = GetAllX(path, files, 0) //抓取所有文件
	for i := 0; i < len(files); i++ {  //打印路径
		fmt.Println(files[i])
	}
}

//
//--
//----

func GetAllX(path string, files []string, level int) ([]string, error) {
	levelstr := ""
	if level == 1 {
		levelstr = "+"
	} else {
		for ; level > 1; level-- {
			levelstr += "|--"
		}
		levelstr += "+"
	}
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件夹不可读取")
	}
	for _, fi := range read { //循环每个文件或文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "\\" + fi.Name()          //构造新的路径
			files = append(files, levelstr+fulldir)     //追加路径
			files, _ = GetAllX(fulldir, files, level+1) //文件夹递归处理
		} else {
			fulldir := path + "\\" + fi.Name()      //构造新的路径
			files = append(files, levelstr+fulldir) //追加路径
		}
	}
	return files, nil
}
