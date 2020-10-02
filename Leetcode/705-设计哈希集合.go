package Leetcode

type MyHashSet struct {
	bucketArr [769]*bucket
}

//map底层实现即为数组
type bucket struct {
	container []int
}

/** Initialize your data structure here. */
func HashConstructor() MyHashSet {
	bucketlist := [769]*bucket{}
	for i := 0; i < 769; i++ {
		bucketlist[i] = new(bucket)
	}
	return MyHashSet{
		bucketArr: bucketlist,
	}
}

func (ms *MyHashSet) Add(key int) {
	ms.bucketArr[ms._hash(key)].insert(key)
}

func (ms *MyHashSet) Remove(key int) {
	ms.bucketArr[ms._hash(key)].delete(key)
}

/** Returns true if this set contains the specified element */
func (ms *MyHashSet) Contains(key int) bool {
	return ms.bucketArr[ms._hash(key)].exists(key)
}

func (ms *MyHashSet) _hash(key int) int {
	return key % 769
}

func (b *bucket) insert(key int) {
	for _, value := range b.container {
		if value == key {
			return
		}
	}
	b.container = append(b.container, key)
}

func (b *bucket) delete(key int) {
	for index, value := range b.container {
		if value == key {
			b.container = append(b.container[0:index], b.container[index+1:]...)
			return
		}
	}
}

func (b *bucket) exists(key int) bool {
	for _, value := range b.container {
		if value == key {
			return true
		}
	}
	return false
}
