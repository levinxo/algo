package main

import (
    "fmt"
)

/**
 * 插入排序
 * Time Complexity O(n^2)
 * Space Complexity 总共O(n)，需要辅助空间O(1)
 */
func InsertSort(data []int){
    length := len(data)
    if length < 2 {
        return
    }

    for i:=1; i<length; i++ {
        for j:=i-1; j>=0; j-- {
            if data[j] > data[j+1] {
                data[j], data[j+1] = data[j+1], data[j]
            } else {
                break;
            }
        }
    }
}

func main() {
    data := []int{22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, -1, 70}
    fmt.Println(data)
    InsertSort(data)
    fmt.Println(data)
}

