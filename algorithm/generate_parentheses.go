package main

import "fmt"

//括号生成
//leetcode: https://leetcode.com/problems/generate-parentheses/

//利用递归+剪枝
//2*n就是我们最后生成字符串的长度，每一个字符可能是"("或")"
//所以，一共有2的2*n次幂的字符串
//递归用于生成所有可能的字符串，加上剪枝限制条件来减少递归次数（即有限制的递归）
//Time complexity 比O(4^n)要少
//Space complexity 比O(4^n)要少

var res []string

func generateParenthesis(n int) []string{
    current := ""
    res = []string{}
    helper(current, n, n)
    return res
}

func helper(current string, left int, right int) {
    if left == 0 && right == 0 {
        res = append(res, current)
        return
    }

    if left > 0 {
        helper(current + "(", left-1, right)
    }
    if right > left {
        helper(current + ")", left, right-1)
    }
}

func main(){
    n := 2
    ret := generateParenthesis(n)
    fmt.Println("output:")
    for _, val := range ret {
        fmt.Println(val)
    }
}

