package main

import "fmt"

//交换单链表结点对
//leetcode: https://leetcode.com/problems/swap-nodes-in-pairs/

type ListNode struct {
    Val int
    Next *ListNode
}

//Iterative
//Time complexity O(n)
//Space complexity O(1)
func swapPairs1(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    var prev  = &ListNode{}
    prev.Next = head
    first := head.Next

    //两个结点为一组，当前后两个结点都存在才进行比较
    for prev.Next != nil && prev.Next.Next != nil {
        a, b := prev.Next, prev.Next.Next
        prev.Next, b.Next, a.Next = b, a, b.Next
        prev = a
    }
    return first
}

//Recursive
//Time complexity O(n)
//Space complexity O(n)
func swapPairs2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    //两个结点为一组进行递归
    blockHead := head.Next
    head.Next = swapPairs2(head.Next.Next)
    blockHead.Next = head

    return blockHead
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    }
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

    head := swapPairs1(&l1)
    printList(head)
    head = swapPairs2(head)
    printList(head)
}

