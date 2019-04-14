package main

import "fmt"

//单链表是否成环
//leetcode: https://leetcode.com/problems/linked-list-cycle/

type ListNode struct {
    Val int
    Next *ListNode
}

//Iterative
//Time complexity O(n)
//Space complexity O(1)
func hasCycle1(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    for step1, step2 := head.Next, head.Next.Next;
    step1 != nil && step2 != nil && step2.Next != nil;
    step1, step2 = step1.Next, step2.Next.Next {
        fmt.Printf("Round: step1-> %d\tstep2-> %d\n", step1.Val, step2.Val)
        if step1 == step2 {
            return true
        }
    }
    return false
}

//Map
//Time complexity O(n)
//Space complexity O(n)
func hasCycle2(head *ListNode) bool {

    for p, hashMap := head, make(map [string] *ListNode);
    p != nil && p.Next != nil; p = p.Next {
        fmt.Printf("Round val: %d\n", p.Val)
        key := fmt.Sprintf("%p", p)
        if _, ok := hashMap[key]; ok {
            return true
        }
        hashMap[key] = p
    }
    return false
}

func main(){
    l1 := ListNode{}
    l1.Val = 1
    l2 := ListNode{}
    l2.Val = 2
    l3 := ListNode{}
    l3.Val = 3
    l4 := ListNode{}
    l4.Val = 4
    l5 := ListNode{}
    l5.Val = 5
    l6 := ListNode{}
    l6.Val = 6

    l1.Next = &l2
    l2.Next = &l3
    l3.Next = &l4
    l4.Next = &l5
    l5.Next = &l6
    l6.Next = &l1

    head := &l1
    if hasCycle1(head) {
        fmt.Println("hasCycle1: has cycle")
    } else {
        fmt.Println("hasCycle1: no cycle")
    }
    if hasCycle2(head) {
        fmt.Println("hasCycle2: has cycle")
    } else {
        fmt.Println("hasCycle2: no cycle")
    }
}

