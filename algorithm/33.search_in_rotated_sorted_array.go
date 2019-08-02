package main

import "fmt"

//33. 搜索旋转排序数组
//https://leetcode.com/problems/search-in-rotated-sorted-array/

//将数组一份为二，其中有个是有序的，另一个是部分有序，
//若target在有序的数组里，就继续执行二分法
//否则target可能在部分有序的数组里，继续寻找下去。
//Time Complexity O(log(n))
//Auxiliary Space Complexity O(1)
func search(nums []int, target int) int {
    return helper(nums, 0, len(nums)-1, target)
}

func helper(nums []int, start int, end int, target int) int {
    if start > end {
        return -1
    }

    mid := start + (end - start) / 2
    if nums[mid] == target {
        return mid
    }
    if nums[mid] < nums[end] {
        if nums[mid] < target && target <= nums[end] {
            return helper(nums, mid+1, end, target)
        } else {
            return helper(nums, start, mid-1, target)
        }
    } else {
        if nums[start] <= target && target < nums[mid] {
            return helper(nums, start, mid-1, target)
        } else {
            return helper(nums, mid+1, end, target)
        }
    }
}

func main(){
    arr := []int{4,5,6,7,0,1,2}
    fmt.Println(search(arr, 0))
    fmt.Println(search(arr, 3))
}

