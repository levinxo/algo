package main

import "fmt"

//二叉树的三种遍历方式
//若中序遍历二叉搜索树，可得到一个顺序数组

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func traversalPreorder(root *TreeNode){
    if root == nil {
        return
    }
    fmt.Printf("root: %d\n", root.Val)
    traversalPreorder(root.Left)
    traversalPreorder(root.Right)
}

func traversalInorder(root *TreeNode){
    if root == nil {
        return
    }
    traversalInorder(root.Left)
    fmt.Printf("root: %d\n", root.Val)
    traversalInorder(root.Right)
}

func traversalPostorder(root *TreeNode){
    if root == nil {
        return
    }
    traversalPostorder(root.Left)
    traversalPostorder(root.Right)
    fmt.Printf("root: %d\n", root.Val)
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

    fmt.Println("pre-order")
    traversalPreorder(&a1)
    fmt.Println("in-order")
    traversalInorder(&a1)
    fmt.Println("post-order")
    traversalPostorder(&a1)
}

