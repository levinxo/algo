package main

import "fmt"

//两数之和
//leetcode: https://leetcode.com/problems/two-sum/

//使用hashmap。假设y变量为target-x，并看这个变量是否存在map中
//Time complexity O(n)
//Space complexity O(n)
func twoSum(nums []int, target int) []int {
    hashTable := make(map[int] int)
    for x_index, x := range nums {
        y := target - x
        if y_index, ok := hashTable[y]; ok && y_index != x_index {
            return []int{x_index,y_index}
        }
        hashTable[x] = x_index
    }
    return []int{}
}

func main(){
    nums := []int{2, 7, 11, 15}
    target := 9
    ret := twoSum(nums, target)
    fmt.Println(ret)
}

