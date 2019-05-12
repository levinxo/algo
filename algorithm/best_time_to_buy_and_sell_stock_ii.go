package main

import "fmt"

//买卖股票的最佳时机 II
//https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

//@todo 还有动态规划

//贪心算法，在每一步都求当前最优解
//Time Complexity O(n)
//Space Complexity O(1)
func maxProfit1(prices []int) int {
    profit := 0
    for i:=1; i<len(prices); i++ {
        if prices[i] > prices[i-1] {
            profit += prices[i] - prices[i-1]
        }
    }
    return profit
}

//按照周期波动计算，从每个波谷到每个波峰
//Time Complexity O(n)
//Space Complexity O(1)
func maxProfit2(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    min, max := prices[0], prices[0]
    profit, i := 0, 0

    for i<len(prices)-1 {
        for i<len(prices)-1 && prices[i] > prices[i+1] {
            i++
        }
        min = prices[i]
        for i<len(prices)-1 && prices[i] <= prices[i+1] {
            i++
        }
        max = prices[i]
        profit += max - min
    }
    return profit
}

func main(){
    a := []int{7, 1, 5, 3, 6, 4}
    //a := []int{1,2,3,4,5}
    //a := []int{7,6,4,3,1}

    fmt.Println("贪心算法：")
    ret := maxProfit1(a)
    fmt.Println(ret)
    fmt.Println("周期波动：")
    ret = maxProfit2(a)
    fmt.Println(ret)
}

