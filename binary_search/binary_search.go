package main

import "fmt"

//二分搜索算法

//递归法
func BinarySearch1(arr []int, start int, end int, target int) int {
    if start > end {
        return -1
    }

    mid := start + (end-start)/2
    if target > arr[mid] {
        return BinarySearch1(arr, mid+1, end, target)
    } else if target < arr[mid] {
        return BinarySearch1(arr, start, mid-1, target)
    } else {
        return mid
    }
}

//迭代法
func BinarySearch2(arr []int, target int) int {
    ret := -1
    if len(arr) == 0 {
        return ret
    }
    start, end := 0, len(arr)-1

    for start <= end {
        mid := start + (end-start)/2
        if target > arr[mid] {
            start = mid+1
        } else if target < arr[mid] {
            end = mid-1
        } else {
            ret = mid
            break
        }
    }
    return ret
}

func main(){
    arr := []int{1,2,5,8,11,23,45,67,78,100,101}
    target := 2
    pos := BinarySearch1(arr, 0, len(arr)-1, target)
    fmt.Println(pos)

    pos = BinarySearch2(arr, target)
    fmt.Println(pos)
}

