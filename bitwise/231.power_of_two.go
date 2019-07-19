package main

import "fmt"

//2的幂 https://leetcode.com/problems/power-of-two/
//解题思路：
//1.取2的模
//2.log函数
//3.位运算，2的幂的二进制一般长这样：0000100000，即仅有一个1

//取2的模
func isPowerOfTwo1(n int) bool {
    if n == 1 || n == 2 {
        return true
    }

    for n > 2 {
        //如果最低位为1，可以直接return
        if n&1 == 1 {
            return false
        }
        //向右移一位，若等于2，则肯定可以被2整除
        n = n>>1
        if n == 2 {
            return true
        }
    }
    return false
}

//位运算
func isPowerOfTwo3(n int) bool {
    if n != 0 && n&(n-1) == 0 {
        return true
    }
    return false
}

func main(){
    nums := [...]int{0,1,2,3,4,8,9,100,200,1024}
    for _, n := range nums {
        fmt.Println(n, ": ", isPowerOfTwo1(n))
    }
    for _, n := range nums {
        fmt.Println(n, ": ", isPowerOfTwo3(n))
    }
}
