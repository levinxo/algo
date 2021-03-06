package main

import "fmt"

//二叉树的最近公共祖先
//leetcode: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//@todo 还有迭代法

//递归法
//如果两边都能找到，那么说明当前root即最近的父节点
//Time complexity O(n)
//Space complexity O(n)
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return root
    }

    //如果找到两个结点中任意一个，就返回
    if root == p || root == q {
        return root
    }

    left := lowestCommonAncestor1(root.Left, p, q)
    right := lowestCommonAncestor1(root.Right, p, q)

    if left == nil {
        return right
    }
    if right == nil {
        return left
    }
    return root
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

    fmt.Println("递归法")
    ret := lowestCommonAncestor1(&a1, &a4, &a5)
    fmt.Println(ret.Val)
}

