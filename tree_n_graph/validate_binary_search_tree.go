package main

import "fmt"

//验证二叉搜索树
//leetcode: https://leetcode.com/problems/validate-binary-search-tree/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//@todo 还有另外两种解法，一种是中序遍历，一种是普通的递归方法

//每个节点都有一个最大最小值限定
//我们从头节点开始进行递归判断即可
//Time complexity O(n)
//Space complexity O(log n)
func isValidBST1(root *TreeNode) bool {
    //定义两个极限最大最小值
    //使用位运算，先对unsigned int取反后向右移一位，
    //就可获得最大的signed int，再取反可获得最小的signed int
    max := int(^uint(0) >> 1)
    min := ^max

    if root == nil {
        return true
    }
    return helper(root.Left, root.Val, min) && helper(root.Right, max, root.Val)
}

func helper(root *TreeNode, max, min int) bool {
    if root == nil {
        return true
    }
    if root.Val >= max || root.Val <= min {
        return false
    }
    fmt.Printf("root: %d, max: %d, min: %d\n", root.Val, max, min)
    return helper(root.Left, root.Val, min) && helper(root.Right, max, root.Val)
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

    ret := isValidBST1(&a1)
    fmt.Println(ret)
}

