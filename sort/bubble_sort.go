package main

import (
    "fmt"
)

/**
 * 冒泡排序
 * Time Complexity O(n^2)
 * Space Complexity 总共O(n)，辅助空间O(1)
 */
func BubbleSort(data []int){
    length := len(data)

    for i:=0; i<length-1; i++ {
        isChange := false
        for j:=0; j<length-1-i; j++ {
            if data[j] > data[j+1] {
                data[j], data[j+1] = data[j+1], data[j]
                isChange = true
            }
        }
        if isChange == false {
            break
        }
    }
}

func main() {
    data := []int{22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, -1, 70}
    fmt.Println(data)
    BubbleSort(data)
    fmt.Println(data)
}

