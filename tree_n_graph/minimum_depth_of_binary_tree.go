package main

import "fmt"

//二叉树的最小深度
//leetcode: https://leetcode.com/problems/minimum-depth-of-binary-tree/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//广度优先搜索
//利用队列把左右结点放入队尾
//Time Complexity O(n)
//Space Complexity O(n)
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    queue, level := []*TreeNode{}, 0

    queue = append(queue, root)
    for len(queue) > 0 {
        level++

        length := len(queue)
        for i:=0; i<length; i++ {
            root = queue[0]
            queue = queue[1:]

            if root.Left == nil && root.Right == nil {
                return level
            }
            if root.Left != nil {
                queue = append(queue, root.Left)
            }
            if root.Right != nil {
                queue = append(queue, root.Right)
            }
        }
    }
    return level
}

func main(){
    /**
    * BST
    *              a1                           10
    *            /    \                      /     \
    *           a2     a3                   5       15
    *         /  \     /  \               /  \     /  \
    *       a4   a5   a6    a7           3    7   12    18
    *      /\   /\    / \   / \         /\   /\   /\    /\
    *    a8 a9 b1 b2 b3 b4 b5 b6       1  4 6  8 11 13 16 19
    * 
    */

    a1 := TreeNode{}; a2 := TreeNode{}; a3 := TreeNode{}
    a4 := TreeNode{}; a5 := TreeNode{}; a6 := TreeNode{}
    a7 := TreeNode{}; a8 := TreeNode{}; a9 := TreeNode{}
    b1 := TreeNode{}; b2 := TreeNode{}; b3 := TreeNode{}
    b4 := TreeNode{}; b5 := TreeNode{}; b6 := TreeNode{}

    a1.Val = 10; a2.Val = 5;  a3.Val = 15; a4.Val = 3
    a5.Val = 7;  a6.Val = 12; a7.Val = 18; a8.Val = 1
    a9.Val = 4;  b1.Val = 6;  b2.Val = 8;  b3.Val = 11
    b4.Val = 13; b5.Val = 16; b6.Val = 19

    a1.Left  = &a2; a1.Right = &a3; a2.Left  = &a4
    a2.Right = &a5; a3.Left  = &a6; a3.Right = &a7
    a4.Left  = &a8; a4.Right = &a9; a5.Left  = &b1
    a5.Right = &b2; a6.Left  = &b3; a6.Right = &b4
    a7.Left  = &b5; a7.Right = &b6

    fmt.Println("广度优先搜索，对每层batch process")
    ret := minDepth(&a1)
    fmt.Println(ret)
}

