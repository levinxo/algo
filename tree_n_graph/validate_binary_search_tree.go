package main

import "fmt"

//验证二叉搜索树
//leetcode: https://leetcode.com/problems/validate-binary-search-tree/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//迭代法，使用stack，深度优先搜索（preorder），也可以借助队列实现广度优先搜索
//Time complexity O(n)
//Space complexity O(n)
func isValidBST2(root *TreeNode) bool {
    max := int(^uint(0) >> 1)
    min := ^max

    var stack []*TreeNode
    var maxStack, minStack []int

    var push = func (r *TreeNode, u, l int) {
        stack = append(stack, r)
        maxStack = append(maxStack, u)
        minStack = append(minStack, l)
    }
    push(root, max, min)

    for len(stack) > 0 {
        root := stack[len(stack)-1]
        max = maxStack[len(maxStack)-1]
        min = minStack[len(minStack)-1]

        stack = stack[:len(stack)-1]
        maxStack = maxStack[:len(maxStack)-1]
        minStack = minStack[:len(minStack)-1]

        if root == nil {
            continue
        }

        val := root.Val
        fmt.Printf("root: %d, max: %d, min: %d\n", val, max, min)
        if val >= max || val <= min {
            return false
        }
        push(root.Right, max, val)
        push(root.Left, val, min)
    }
    return true
}

//递归法，每个节点都有一个最大最小值限定
//我们从头节点开始进行递归判断即可
//Time complexity O(n)
//Space complexity O(n)
func isValidBST1(root *TreeNode) bool {
    //定义两个极限最大最小值
    //使用位运算，先对unsigned int取反后向右移一位，
    //就可获得最大的signed int，再取反可获得最小的signed int
    //或者设置为nil也是可以的，更合适一些
    max := int(^uint(0) >> 1)
    min := ^max

    return helper(root, max, min)
}

func helper(root *TreeNode, max, min int) bool {
    if root == nil {
        return true
    }
    fmt.Printf("root: %d, max: %d, min: %d\n", root.Val, max, min)
    if root.Val >= max || root.Val <= min {
        return false
    }
    return helper(root.Left, root.Val, min) && helper(root.Right, max, root.Val)
}

//中序遍历，深度优先搜索(inorder)
//Time complexity O(n)
//Space complexity O(n)
func isValidBST3(root *TreeNode) bool {
    var stack []*TreeNode
    min := ^int(^uint(0)>>1)

    for len(stack) > 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }

        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        fmt.Printf("root: %d, min: %d\n", root.Val, min)
        if root.Val > min {
            min = root.Val
            root = root.Right
        } else {
            return false
        }
    }
    return true
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
    ret := isValidBST1(&a1)
    fmt.Println(ret)
    fmt.Println("迭代深搜")
    ret = isValidBST2(&a1)
    fmt.Println(ret)
    fmt.Println("深搜中序遍历")
    ret = isValidBST3(&a1)
    fmt.Println(ret)
}

