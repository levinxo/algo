package main

import "fmt"

//二叉树的层次遍历
//leetcode: https://leetcode.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

//广度优先搜索
//利用队列把左右结点放入队尾
//Time Complexity O(n)
//Space Complexity O(n)
func levelOrder1(root *TreeNode) [][]int {

    row, res, queue := []int{}, [][]int{}, []*TreeNode{}
    thisRowRemain, nextRowHas := 1, 0

    if root == nil {
        return res
    }

    queue = append(queue, root)

    for len(queue) > 0 {
        //pop
        root = queue[0]
        queue = queue[1:]
        thisRowRemain--

        row = append(row, root.Val)

        if root.Left != nil {
            queue = append(queue, root.Left)
            nextRowHas++
        }
        if root.Right != nil {
            queue = append(queue, root.Right)
            nextRowHas++
        }

        if thisRowRemain == 0 {
            res = append(res, row)
            row = []int{}
            thisRowRemain, nextRowHas = nextRowHas, 0
        }
    }
    return res
}

//广度优先搜索，batch process
//利用队列把左右结点放入队尾
//Time Complexity O(n)
//Space Complexity O(n)
func levelOrder2(root *TreeNode) [][]int {
    row, res, queue := []int{}, [][]int{}, []*TreeNode{}

    if root == nil {
        return res
    }

    queue = append(queue, root)

    for len(queue) > 0 {

        length := len(queue)
        for i:=0; i<length; i++ {
            root = queue[0]
            queue = queue[1:]
            row = append(row, root.Val)
            if root.Left != nil {
                queue = append(queue, root.Left)
            }
            if root.Right != nil {
                queue = append(queue, root.Right)
            }
        }
        res = append(res, row)
        row = []int{}
    }
    return res
}

//深度优先搜索，preorder
//利用递归进行深度搜索，并借助level透传来标记层级
//Time Complexity O(n)
//Space Complexity O(n)
func levelOrder3(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    res := &([][]int{})
    res = helper(root, 0, res)
    return *res
}

func helper(root *TreeNode, level int, res *[][]int) *[][]int {
    if level > len(*res)-1 {
        *res = append(*res, []int{root.Val})
    } else {
        (*res)[level] = append((*res)[level], root.Val)
    }

    if root.Left != nil {
        res = helper(root.Left, level+1, res)
    }
    if root.Right != nil {
        res = helper(root.Right, level+1, res)
    }
    return res
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

    fmt.Println("广度优先搜索，一次循环")
    ret := levelOrder1(&a1)
    fmt.Println(ret)
    fmt.Println("广度优先搜索，对每层batch process")
    ret = levelOrder2(&a1)
    fmt.Println(ret)
    fmt.Println("深度优先搜索")
    ret = levelOrder3(&a1)
    fmt.Println(ret)
}

