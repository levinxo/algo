package main

import "fmt"

//找出单链表环开始结点
//leetcode: https://leetcode.com/problems/linked-list-cycle-ii/

type ListNode struct {
    Val int
    Next *ListNode
}

//Iterative
//Time complexity O(n)
//Space complexity O(1)
//
//        ____
//       |    a2
//o-----a1    |
//       |    |
//       |_a3_|
//
//两个指针，p1速度为1，p2速度为2；o为起点；环为顺时针方向
//p1从o到a1时走过的路程为X，则此时p2走到a2，且a2-a1=X（因为p2速度是p1的两倍）
//设环的周长为C，此时p2若想追上p1，他们距离差C-X，速度差1，则需要时间为t=(C-X)/1。
//在时间t内，p1走过的路程为C-X，到达a3，则a3至a1的距离为X
//已知o到a1的距离为X。所以o至a1和a3至a1距离相等。
//因此此时再把p2放回o点，速度保持为1，p1和p2再次相遇时，即为环的入口a1。
//
func detectCycle1(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    for p1,p2:=head.Next,head.Next.Next;
    p1 != nil && p2 != nil && p2.Next != nil;
    p1, p2 = p1.Next, p2.Next.Next {
        fmt.Printf("Loop1.Round: p1-> %d\tp2-> %d\n", p1.Val, p2.Val)

        //相遇时，把p2移回链表头
        if p1 == p2 {
            for p2 = head; ; p1,p2=p1.Next,p2.Next {
                fmt.Printf("Loop2.Round: p1-> %d\tp2-> %d\n", p1.Val, p2.Val)
                if p1 == p2 {
                    return p1
                }
                //若有空变量，就退出
                if p1 == nil || p2 == nil || p1.Next == nil || p2.Next == nil {
                    return nil
                }
            }
        }
    }
    return nil
}

//Map
//Time complexity O(n)
//Space complexity O(n)
func detectCycle2(head *ListNode) *ListNode {
    for hashMap, p := make(map [string] *ListNode), head;
    p != nil && p.Next != nil; p = p.Next {
        key := fmt.Sprintf("%p", p)
        fmt.Printf("Round: p-> %d\n", p.Val)
        if _, ok := hashMap[key]; ok {
            return p
        }
        hashMap[key] = p
    }
    return nil
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
    p := detectCycle1(head)
    if p != nil {
        fmt.Printf("cycle begin at %d\n", p.Val);
    } else {
        fmt.Printf("no cycle\n")
    }
    p = detectCycle2(head)
    if p != nil {
        fmt.Printf("cycle begin at %d\n", p.Val);
    } else {
        fmt.Printf("no cycle\n")
    }
}
