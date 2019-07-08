package main

import (
    "fmt"
)

/**
 * 冒泡排序
 * Time Complexity O(n^2)
 * Space Complexity 总共O(n)，辅助空间O(1)
 */
func BubbleSort(arr []int){
    length := len(arr)

    for i:=0; i<length-1; i++ {
        isChange := false
        for j:=0; j<length-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
                isChange = true
            }
        }
        if isChange == false {
            break
        }
    }
}

func main() {
    arr := []int{22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, -1, 70}
    fmt.Println(arr)
    BubbleSort(arr)
    fmt.Println(arr)
}

