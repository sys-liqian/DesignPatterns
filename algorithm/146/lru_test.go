package a46

import (
	"container/list"
	"testing"
)

//LRUCache(int capacity) 以正整数作为容量capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value)如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
//当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

func TestLru(t *testing.T) {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	M   map[int]int
	Cap int
	L   *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		M:   make(map[int]int, capacity),
		Cap: capacity,
		L:   list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	value := c.M[key]
	if value >= 0 {
		//让被访问的元素放在链表头部
		c.L.PushFront(key)
		return value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	mapSize := len(c.M)
	//put时判断LRU当前的容量
	if mapSize >= c.Cap {
		//移除最近不使用的元素

	}
	//插入元素
	c.M[key] = value

}
