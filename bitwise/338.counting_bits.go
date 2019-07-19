package main

import "fmt"

//https://leetcode.com/problems/counting-bits/
//1.循环数组，每个数都去计算它的汉明重量
//2.递推，利用去掉低位1的数的汉明重量+1来得到当前数的汉明重量

func countBits(num int) []int {
    res := make([]int, num+1)

    for i:=1; i<=num; i++ {
        res[i] = res[i&(i-1)] + 1
    }

    return res
}

func main(){
    num := 100
    fmt.Println(countBits(num))
}
