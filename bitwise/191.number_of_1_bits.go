package main

import "fmt"

//191. 位1的个数 https://leetcode.com/problems/number-of-1-bits/
//两种方法：
//1.取模判断为1时就++，然后对数字进行右移
//2.num和num-1进行与操作，可以去掉最低位的1

func hammingWeight1(num uint32) int {
    ret := 0
    for num > 0 {
        if num%2 == 1 {
            ret++
        }
        num = num>>1
    }
    return ret
}

func hammingWeight2(num uint32) int {
    ret := 0
    for num != 0 {
        ret++
        num = num&(num-1)
    }
    return ret
}

func main(){
    num := uint32(100)
    count := hammingWeight1(num)
    fmt.Println(count)
    count = hammingWeight1(num)
    fmt.Println(count)
}
