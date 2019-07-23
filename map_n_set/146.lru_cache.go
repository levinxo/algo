package main

import "fmt"

//146. LRU缓存机制 https://leetcode.com/problems/lru-cache/
//
//第一种方法：使用两种数据结构：哈希表 + 双向链表
//结点存了值，哈希表的元素和双向链表都保存了结点的地址
//新的key直接存入链表头部，依靠双向链表达到淘汰的目的
//
//第二种方法：使用有序哈希表，go暂时不支持，需要自己实现

type TreeNode struct {
    key int
    val int
    next *TreeNode
    prev *TreeNode
}

type LRUCache struct {
    Cache map[int] *TreeNode
    HeadNode *TreeNode
    TailNode *TreeNode
    Size int
    Capacity int
}

func Constructor(capacity int) LRUCache {
    max := int(^uint(0)>>1)
    head, tail := &TreeNode{max, 0, nil, nil}, &TreeNode{max, 0, nil, nil}
    head.next = tail
    tail.prev = head
    return LRUCache{
        make(map[int]*TreeNode),
        head,
        tail,
        0,
        capacity,
    }
}

func (this *LRUCache) addNode(node *TreeNode){
    node.prev = this.HeadNode
    node.next = this.HeadNode.next

    this.HeadNode.next.prev = node
    this.HeadNode.next = node
}

func (this *LRUCache) removeNode(node *TreeNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) deleteTail() *TreeNode {
    delNode := this.TailNode.prev
    this.removeNode(delNode)
    this.Size--
    return delNode
}

//Time Complexity O(1)
//Space Complexity O(n)
func (this *LRUCache) Get(key int) int {
    if node, ok := this.Cache[key]; !ok {
        return -1
    } else {
        //放置到双链表头部
        this.removeNode(node)
        this.addNode(node)
        return node.val
    }
}

//Time Complexity O(1)
//Space Complexity O(n)
func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.Cache[key]; !ok {
        node = &TreeNode{key, value, nil, nil}
        this.addNode(node)
        this.Cache[key] = node
        this.Size++
        if this.Size > this.Capacity {
            tmp := this.deleteTail()
            delete(this.Cache, tmp.key)
        }
    } else {
        node.val = value
        this.removeNode(node)
        this.addNode(node)
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
    capacity := 5
    obj := Constructor(capacity)

    fmt.Println(obj.Get(1))
    obj.Put(1, 1)
    fmt.Println(obj.Get(1))
    obj.Put(2, 2)
    fmt.Println(obj.Get(2))
    obj.Put(3, 3)
    fmt.Println(obj.Get(3))
    obj.Put(4, 4)
    fmt.Println(obj.Get(4))
    obj.Put(5, 5)
    fmt.Println(obj.Get(5))

    fmt.Println(obj.Get(1))
    fmt.Println(obj.Get(2))
    fmt.Println(obj.Get(3))
    fmt.Println(obj.Get(4))
    fmt.Println(obj.Get(5))

    obj.Put(6, 6)
    fmt.Println(obj.Get(6))
    fmt.Println(obj.Get(1))
}
