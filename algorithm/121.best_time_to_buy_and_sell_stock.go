package main

import "fmt"

//买卖股票的最佳时机
//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

//维护两个变量：min和maxP
//Time Complexity O(k)
//Auxiliary Space Complexity O(1)
func maxProfit(prices []int) int {
    min := int(^uint(0) >> 1)
    maxP := 0

    for _, price := range prices {
        if price < min {
            min = price
        } else if price - min > maxP {
            maxP = price - min
        }
    }
    return maxP
}

func main(){
    a := []int{7,1,5,3,6,4}
    profit := maxProfit(a)
    fmt.Println(profit)
}

