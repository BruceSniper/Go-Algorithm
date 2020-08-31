package HashTableArray

import (
	"crypto/sha256"
	"errors"
)

const (
	Delete       = iota //数据已经被删除
	MinTableSize = 100  //哈希表的大小
	legimate     = iota //已经存在的合法数据
	Empty        = iota //数据为空

)

//自定义哈希函数
func MySHA(str interface{}, tableSize int) int {
	var Hashvar int = 0
	var chars []byte
	if strings, ok := str.(string); ok {
		chars = []byte(strings) //字符串转化为字节数
	}
	for _, v := range chars {
		Hashvar = (Hashvar << 17) + int(v) //哈希算法
	}
	return Hashvar % MinTableSize
}

func MySHA256(str string, tableSize int) int {
	var Hashvar int = 0
	shaobj := sha256.New()
	shaobj.Write([]byte(str)) //哈希
	mybytes := shaobj.Sum(nil)
	for _, v := range mybytes {
		Hashvar = (Hashvar<<17 | 123&1235 ^ 139) + int(v) //哈希算法
	}
	return Hashvar % MinTableSize
}

type HashFunc func(data interface{}, tableSize int) int //函数指针

type HashEntry struct {
	data interface{} //数据
	kind int         //类型

}

type HashTable struct {
	tableSize int          //哈希表的大小
	theCells  []*HashEntry //数组，每一个元素是指针指向哈希结构
	hashfunc  HashFunc     //调用哈希函数
}

type HashTableer interface {
	Find(data interface{}) int      //查找数据
	Insert(data interface{})        //插入数据
	Empty()                         //为空
	GetValue(index int) interface{} //抓取value
}

func NewHashTable(size int, hash HashFunc) (*HashTable, error) {
	if size < MinTableSize {
		return nil, errors.New("哈希表太小")
	}
	if hash == nil {
		return nil, errors.New("没有哈希函数")
	}
	hashtable := new(HashTable)                   //创建哈希表
	hashtable.tableSize = size                    //设置哈希表的大小
	hashtable.theCells = make([]*HashEntry, size) //数组分配内存
	hashtable.hashfunc = hash                     //设置哈希函数
	for i := 0; i < hashtable.tableSize; i++ {
		hashtable.theCells[i] = new(HashEntry)
		hashtable.theCells[i].data = nil
		hashtable.theCells[i].kind = Empty //设置为空
	}
	return hashtable, nil
}

func (ht *HashTable) Find(data interface{}) int {
	var collid int = 0
	curpos := ht.hashfunc(data, ht.tableSize) //计算哈希位置
	if ht.theCells[curpos].kind != Empty && ht.theCells[curpos].data != data {
		collid += 1
		curpos = 2*curpos - 1 // 平方探测
		if curpos > ht.tableSize {
			curpos -= ht.tableSize //越界，返回
		}
	}
	return curpos
}

func (ht *HashTable) Insert(data interface{}) {
	pos := ht.Find(data)      //查找数据位置
	entry := ht.theCells[pos] //插入数据记录状态
	if entry.kind != legimate {
		entry.kind = legimate
		entry.data = data //插入数据
	}
}

func (ht *HashTable) Empty() {
	for i := 0; i < ht.tableSize; i++ {
		if ht.theCells[i] == nil {
			continue //循环清空数据
		}
		ht.theCells[i].kind = Delete //删除数据
	}
}

func (ht *HashTable) GetValue(index int) interface{} {
	if index > ht.tableSize {
		return nil //判断大小
	}
	entry := ht.theCells[index] //取出数据
	if entry.kind == legimate {
		return entry.data
	} else {
		return nil
	}
}
