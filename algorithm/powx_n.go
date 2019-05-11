package main

import "fmt"

//Pow(x, n)
//https://leetcode.com/problems/powx-n/description/

//利用递归的分治法
//Time complexity O(log n)
//Space complexity O(?)
func myPow1(x float64, n int) float64 {

    fmt.Printf("param n: %d\tparam x: %f\n", n, x)
    if n == 0 {
        return 1
    }
    if n < 0 {
        return 1/myPow1(x, -n)
    }
    if n % 2 != 0 {
        return x * myPow1(x*x, (n-1)/2)
    }
    return myPow1(x*x, n/2)
}

//利用迭代的分治法
//Time complexity O(log n)
//Space complexity O(1)
func myPow2(x float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n < 0 {
        x = 1/x
        n = -n
    }

    var res float64
    res = 1

    for n > 0 {
        fmt.Printf("n=%d, ", n)
        //奇数时，这里res把n为奇数的时候多余的x给留存起来，
        //并在最后n=1的时候乘以x最终值，再return
        if n&1 == 1 {
            fmt.Printf("res=%f, x=%f, ", res, x)
            res *= x
        }
        fmt.Printf("x=%f\n",  x)
        x *= x
        n >>= 1
    }
    return res
}

func main(){
    var (
        x float64
        n int
    )

    x = 2
    n = 5

    ret := myPow1(x, n)
    fmt.Println(ret)
    ret = myPow2(x, n)
    fmt.Println(ret)
}

