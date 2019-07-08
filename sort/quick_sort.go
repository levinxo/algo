package main

import (
    "fmt"
)

/**
 * 快速排序
 * Time Complexity O(n logn)
 * Space Complexity
 */
func QuickSort(data []int){
    if len(data) < 2 {
        return
    }

    pivot := data[0]
    head, tail := 0, len(data) - 1

    for i:=1; i<=tail; {
        if pivot < data[i] {
            data[i], data[tail] = data[tail], data[i]
            tail--
        } else {
            data[i], data[head] = data[head], data[i]
            head++
            i++
        }
    }
    QuickSort(data[:head])
    QuickSort(data[head+1:])
}

func main() {
    data := []int{22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, -1, 70}
    fmt.Println(data)
    QuickSort(data)
    fmt.Println(data)
}

