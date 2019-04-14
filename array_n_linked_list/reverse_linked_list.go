package main

import "fmt"

//反转单链表
//leetcode: https://leetcode.com/problems/reverse-linked-list/

type ListNode struct {
    Val int
    Next *ListNode
}

//Iterative
//Time complexity O(n)
//Space complexity O(1)
//
//①       1->2->3
//②  nil<-1->2->3
//③  nil<-1<-2->3
func reverseList1(head *ListNode) *ListNode {
    var prev,curr *ListNode
    curr = head
    for curr != nil {
        nextTemp := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextTemp
    }
    return prev
}

//Recursive
//Time complexity O(n)
//Space complexity O(n)
//
//①  1->2->3
//
//     nil
//      ↑
//②  1->2<-3
//③  nil<-1<-2<-3
func reverseList2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    //p透传，最终为最后一个结点
    p := reverseList2(head.Next)
    //head下个结点的next指向head
    head.Next.Next = head
    head.Next = nil
    return p
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

    l1.Next = &l2
    l2.Next = &l3

    head := reverseList1(&l1)
    printList(head)
    head = reverseList2(head)
    printList(head)
}

