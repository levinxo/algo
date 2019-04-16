package main

import "container/heap"
import "fmt"

//数据流中的第K大元素
//leetcode: https://leetcode.com/problems/kth-largest-element-in-a-stream/
//Use min-heap

type IntHeap []int

func (h IntHeap) Len() int {return len(h)}

func (h IntHeap) Less(i, j int) bool {return h[i]<h[j]}

func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func (h *IntHeap) Peek() interface{} {
    return (*h)[0]
}

type KthLargest struct {
    K int
    H *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
    max := &IntHeap{}

    heap.Init(max)
    kth := KthLargest{
        K: k,
        H: max,
    }

    for i:=0; i<len(nums); i++ {
        kth.Add(nums[i])
    }
    return kth
}

func (this *KthLargest) Add(val int) int {

    if this.H.Len() < this.K {
        heap.Push(this.H, val)
    } else if this.H.Peek().(int) < val {
        heap.Pop(this.H)
        heap.Push(this.H, val)
    }

    x := this.H.Peek().(int)
    fmt.Printf("peek: %d\n", x)
    return x
}

func main(){
    k := 3
    nums := []int{4,5,8,2}
    obj := Constructor(k, nums);
    obj.Add(10);
    obj.Add(15);
    obj.Add(6);
    obj.Add(7);
}

