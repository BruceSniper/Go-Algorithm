package ArrayList

import (
	"errors"
	"fmt"
)

//接口
type List interface {
	Size() int                                  //数组大小
	Get(index int) (interface{}, error)         //抓取第几个元素
	Set(index int, newval interface{}) error    //修改数据
	Insert(index int, newval interface{}) error //插入数据
	Append(newval interface{})                  //追加数据
	Clear()                                     //清空
	Delete(index int) error                     //删除
	String() string                             //返回字符串
	Iterator() Iterator                         //接口
}

//数据结构，字符串，整数，实数
type ArrayList struct {
	dataStore []interface{} //数组存储
	TheSize   int           //数组的大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)                      //初始化结构体
	list.dataStore = make([]interface{}, 0, 10) //开辟空间10个
	list.TheSize = 0
	return list
}

func (list *ArrayList) checkisFull() {
	if list.TheSize == cap(list.dataStore) { //判断内存使用
		//make 中间的参数如果为0，即没有开辟内存
		newDataStore := make([]interface{}, 2*list.TheSize, 2*list.TheSize) //开辟双倍内存
		copy(newDataStore, list.dataStore)                                  //拷贝
		//或者：
		//for i := 0; i < len(list.dataStore)-1; i++ {
		//	newDataStore[i]=list.dataStore[i]
		//}
		list.dataStore = newDataStore //赋给新的数组
		fmt.Println(list.dataStore)
		fmt.Println(newDataStore)
	}
}

func (list *ArrayList) Size() int {
	return list.TheSize //返回数组的大小
}

//追加函数
func (list *ArrayList) Append(newval interface{}) {
	list.dataStore = append(list.dataStore, newval)
	list.TheSize++
}

//抓取数据
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.TheSize {
		return nil, errors.New("索引越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}

func (list *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.dataStore[index] = newval //设置
	return nil
}

func (list *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= list.TheSize {
		return errors.New("索引越界")
	}
	list.checkisFull()                               //检测内存大小，如果满了自动扩容
	list.dataStore = list.dataStore[:list.TheSize+1] //插入数据，内存移动一位
	for i := list.TheSize; i > index; i-- {          //从后往前移动
		list.dataStore[i] = list.dataStore[i-1]
	}
	list.dataStore[index] = newval //插入数据
	list.TheSize++                 //索引追加

	return nil
}

//清空，重新开辟内存空间即可
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}

func (list *ArrayList) Delete(index int) error { //删除
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...) //重新叠加跳过index
	list.TheSize--

	return nil
}
