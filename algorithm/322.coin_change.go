package main

import "fmt"
import "math"

//leetcode: https://leetcode.com/problems/coin-change/
//322. 零钱兑换

//递归暴力解法
//Time Complexity O(k*n^k) k为硬币种类数
//Space Complexity O(?)
func coinChange1(coins []int, amount int) int {
    max := int(^uint(0) >> 1)
    cnt := max

    if amount == 0 {
        return 0
    }

    for _, coin := range coins {
        if amount - coin  < 0 {
            continue
        }
        nextCnt := coinChange1(coins, amount - coin)
        if nextCnt == -1 {
            continue
        }
        cnt = int(math.Min(float64(cnt), float64(nextCnt+1)))
    }

    if cnt == max {
        cnt = -1
    }
    return cnt
}

//有记忆的递归
//Time Complexity O(k*n)
func coinChange2(coins []int, amount int) int {
    memo := make([]int, amount+1, amount+1)
    for i, _ := range memo {memo[i] = -2}
    return helper(coins, amount, &memo)
}

func helper(coins []int, amount int, memo *[]int) int {
    max := int(^uint(0) >> 1)
    cnt := max

    if amount == 0 {
        return 0
    }
    if (*memo)[amount] != -2 {
        return (*memo)[amount]
    }

    for _, coin := range coins {
        if amount - coin < 0 {
            continue
        }
        nextCnt := helper(coins, amount-coin, memo)
        if nextCnt == -1 {
            continue
        }
        if nextCnt+1 < cnt {
            cnt = nextCnt+1
        }
    }
    if cnt == max {
        cnt = -1
    }
    (*memo)[amount] = cnt
    return (*memo)[amount]
}

//动态规划
//递归的自底向上版本
func coinChange3(coins []int, amount int) int {
    max := int(^uint(0) >> 1)
    dp := make([]int, amount+1, amount+1)
    for i, _ := range dp {dp[i] = max}
    dp[0] = 0

    for i, _ := range dp {
        for _, coin := range coins {
            if i - coin < 0 || dp[i-coin] == max {
                continue
            }
            //状态转移方程
            if 1+dp[i-coin] < dp[i] {
                dp[i] = 1+dp[i-coin]
            }
        }
    }
    if dp[amount] == max {
        return -1
    }
    return dp[amount]
}

func main(){
    coins := []int{1,2,5}
    amount := 11

    fmt.Println(coinChange1(coins, amount))
    fmt.Println(coinChange2(coins, amount))
    fmt.Println(coinChange3(coins, amount))
}
