package main

import (
    "fmt"
)

/**
 * 快速排序
 * Time Complexity O(n logn)
 * Space Complexity
 */
func QuickSort1(data []int){
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
    QuickSort1(data[:head])
    QuickSort1(data[head+1:])
}

//少交换次数的版本
func QuickSort2(data []int, start, end int) {
    if len(data) < 2 || start >= end {
        return
    }

    i, j := start, end
    pivot := data[start + (end - start) / 2]

    for i <= j {
        for data[i] < pivot {
            i++
        }
        for data[j] > pivot {
            j--
        }
        if i < j {
            data[i], data[j] = data[j], data[i]
            i++
            j--
        } else if i == j {
            i++
        }
    }

    QuickSort2(data, start, j)
    QuickSort2(data, i, end)
}

func main() {
    var data []int

    fmt.Println("第一种方法：")
    data = []int{22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, -1, 70}
    QuickSort1(data)
    fmt.Println(data)

    fmt.Println("第二种方法：")
    data = []int{22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, -1, 70}
    QuickSort2(data, 0, len(data)-1)
    fmt.Println(data)
}

