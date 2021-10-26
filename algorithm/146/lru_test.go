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
	lruCache := Constructor(2)

	lruCache.Put(2, 1)

	lruCache.Put(2, 2)
	lruCache.Get(2)

	lruCache.Put(1, 1)
	lruCache.Put(4, 1)

	lruCache.Get(2)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type LRUCache struct {
	M map[int]*list.Element
	C int
	L *list.List
}

type Node struct {
	Key, Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		M: make(map[int]*list.Element, capacity),
		C: capacity,
		L: list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	v := c.M[key]
	if v != nil {
		node := v.Value.(Node)
		//让被访问的元素放在链表头部
		c.L.MoveToFront(v)
		return node.Value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	node := Node{key, value}
	//判断元素是否存在
	if v, ok := c.M[key]; ok {
		v.Value = node
		c.L.MoveToFront(v)
		return
	}
	mapSize := len(c.M)

	//put时判断LRU当前的容量
	if mapSize >= c.C {
		//链表移除最近不使用的元素
		backElement := c.L.Back()
		delKey := backElement.Value.(Node).Key
		//双向列表中删除
		c.L.Remove(backElement)
		//map中删除该key,value
		delete(c.M, delKey)
	}
	//插入元素
	el := c.L.PushFront(node)
	c.M[key] = el
}
