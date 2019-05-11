package main

import "fmt"

//Pow(x, n)
//https://leetcode.com/problems/powx-n/description/

//@todo 还可以用迭代来算

//利用递归的分治法
//Time complexity O(log n)
//Space complexity O(?)
func myPow(x float64, n int) float64 {

    fmt.Printf("param n: %d\tparam x: %f\n", n, x)
    if n == 0 {
        return 1
    }
    if n < 0 {
        return 1/myPow(x, -n)
    }
    if n % 2 != 0 {
        return x * myPow(x*x, (n-1)/2)
    }
    return myPow(x*x, n/2)
}

func main(){
    var (
        x float64
        n int
    )

    x = 2
    n = 4

    ret := myPow(x, n)
    fmt.Println(ret)
}

