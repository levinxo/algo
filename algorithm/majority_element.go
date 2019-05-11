package main

import "fmt"

//求众数
//https://leetcode.com/problems/majority-element/description/

//@todo 利用分治法来做
//@todo 对数组排序后进行遍历计数

//利用hashmap
//Time complexity O(n)
//Space complexity O(n)
func majorityElement(nums []int) int {
    hashTable := make(map[int] int)
    half := len(nums)/2

    for _, x := range nums {
        if _, ok := hashTable[x]; ok {
            hashTable[x]++

            fmt.Println(x, hashTable[x], half)

            if hashTable[x] > half {
                return x
            }
        } else {
            hashTable[x] = 1;

            if hashTable[x] > half {
                return x
            }
        }
    }
    //如果是实际应用中，应该返回两个变量，一个是int，一个是bool
    return 0
}

func main(){
    nums := []int{1,2,2,3,5,2,2}

    ret := majorityElement(nums)
    fmt.Println(ret)
}

