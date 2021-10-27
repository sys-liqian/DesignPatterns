package a46

import (
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
	M map[int]*Node
	L *LinkedList
}

type LinkedList struct {
	Head, Tail     *Node
	Size, Capacity int
}

type Node struct {
	Key, Value int
	Pre, Next  *Node
}

func Constructor(capacity int) LRUCache {
	list := &LinkedList{
		Head:     &Node{},
		Tail:     &Node{},
		Size:     0,
		Capacity: capacity,
	}
	list.Head.Next = list.Tail
	list.Tail.Pre = list.Head
	return LRUCache{
		M: make(map[int]*Node, capacity),
		L: list,
	}
}

func (c *LRUCache) Get(key int) int {
	node := c.M[key]
	if node != nil {
		//让被访问的元素放在链表头部
		c.MoveToFront(node)
		return node.Value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	//判断元素是否存在
	if node, ok := c.M[key]; ok {
		//key存在,将该node移动到链表头部
		c.MoveToFront(node)
		//更新该key下value的值
		c.M[key].Value = value
		return
	}
	//put时判断LRU当前Size和容量
	if c.L.Size >= c.L.Capacity {
		//链表移除最近不使用的元素
		//LinkedList尾节点前一个节点为用户可操作的最后节点
		backElement := c.L.Tail.Pre
		delKey := backElement.Key
		//双向链表中删除
		c.Remove(backElement)
		//map中删除该key,value
		delete(c.M, delKey)
	}
	//插入元素
	putNode := &Node{
		Key:   key,
		Value: value,
	}
	c.PushFront(putNode)
	c.M[key] = putNode
}

func (c *LRUCache) MoveToFront(node *Node) {
	c.Remove(node)
	c.PushFront(node)
}

func (c *LRUCache) PushFront(node *Node) {
	head := c.L.Head

	node.Pre = head
	node.Next = head.Next

	head.Next.Pre = node
	head.Next = node
	c.L.Size += 1
}

func (c *LRUCache) Remove(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	c.L.Size -= 1
}
