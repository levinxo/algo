package main

import (
    "fmt"
)

/**
 * 选择排序
 * Time Complexity O(n^2)
 * Space Complexity O(n) total，O(1) auxiliary
 */
func SelectSort(data []int){
    length := len(data)

    if length < 2 {
        return
    }

    for i:=0; i<length-1; i++ {
        min := i
        for j:=i+1; j<length; j++ {
            if data[min] > data[j] {
                min = j
            }
        }

        if min != i {
            data[min], data[i] = data[i], data[min]
        }
    }
}

func main() {
    data := []int{22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, -1, 70}
    fmt.Println(data)
    SelectSort(data)
    fmt.Println(data)
}

