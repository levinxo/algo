package main

import "fmt"

//买卖股票的最佳时机
//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

//维护两个变量：min和maxP
//Time Complexity O(k)
//Auxiliary Space Complexity O(1)
func maxProfit1(prices []int) int {
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

//动态规划
//状态转移方程：
//dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + price[i])
//dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - price[i])
//
//基础case：
//dp[-1][k][0] = 0
//dp[-1][k][1] = -infinity
//dp[i][0][0] = 0
//dp[i][0][1] = -infinity
func maxProfit2(prices []int) int {
    n := len(prices)
    if n == 0 {
        return 0
    }
    dp := make([][]int, n)
    for i, _ := range dp {dp[i] = make([]int, 2)}

    for i, price := range prices {
        if i == 0 {
            dp[i][0] = 0
            dp[i][1] = -price
            continue
        }
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+price)
        dp[i][1] = max(dp[i-1][1], -price)
    }
    return dp[n-1][0]
}

//动态规划，简化变量
func maxProfit3(prices []int) int {
    n := len(prices)
    if n == 0 {
        return 0
    }
    dp_i_0, dp_i_1 := 0, -int(^uint(0) >> 1)

    for _, price := range prices {
        dp_i_0 = max(dp_i_0, dp_i_1+price)
        dp_i_1 = max(dp_i_1, -price)
    }
    return dp_i_0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main(){
    a := []int{7,1,5,3,6,4}
    //a = []int{7,6,5,4,3,2,1}
    fmt.Println(maxProfit1(a))
    fmt.Println(maxProfit2(a))
    fmt.Println(maxProfit3(a))
}

