package main

import (
    "fmt"
)

/**
 * 归并排序
 * Time Complexity O(n log n)
 * Space Complexity
 */
func MergeSort(data []int) []int {
    length := len(data)
    if length < 2 {
        return data
    }

    key := length / 2

    left := MergeSort(data[:key])
    right := MergeSort(data[key:])
    return merge(left, right)
}

func merge(left, right []int) []int {
    i, j := 0, 0
    tmp := []int{}

    for i<len(left) && j<len(right) {
        if left[i] <= right[j] {
            tmp = append(tmp, left[i])
            i++
        } else {
            tmp = append(tmp, right[j])
            j++
        }
    }

    tmp = append(tmp, left[i:]...)
    tmp = append(tmp, right[j:]...)
    return tmp
}

func main() {
    data := []int{22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, -1, 70}
    fmt.Println(data)
    fmt.Println(MergeSort(data))
}

