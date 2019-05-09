package main

import "fmt"
import "sort"

//三数之和
//leetcode: https://leetcode.com/problems/3sum/

//先遍历x，再遍历y的同时使用hashMap，关键在于如何对y进行去重
//Time complexity O(n^2)
//Space complexity O(n)
func threeSum1(nums []int) [][]int {
    res, target := [][]int{}, 0
    if len(nums) < 3 {
        return res
    }
    //将nums排序，方便后续去重
    sort.Ints(nums)
    //fmt.Println(nums)

    for i, x := range nums[:len(nums)-2] {
        //判重，重复的x没有意义
        if i>0 && nums[i-1] == x {
            continue
        }

        //判重标识，见之后说明
        flag := false
        hashTable := make(map[int]int)

        //从x之后开始遍历y
        for j, y := range nums[i+1:]{
            //if j>0 {
            //    fmt.Printf("%d, %d\n", y, nums[i+j])
            //}
            //通过和前个y对比以及flag标志去重，flag防止误杀
            //在没有生成三个元素的数组之前，是可以允许y重复的
            //i+j，是y的前一位绝对index（j-1是相对index，不用它来定位nums的value）
            if j>0 && nums[i+j] == y && flag {
                continue
            }
            flag = false
            z := target-x-y
            if _, ok := hashTable[z]; ok {
                res = append(res, []int{x,y,z})
                //设置一个flag用来去重，避免y重复
                flag = true
                //fmt.Println([]int{x,y,z})
            } else {
                hashTable[y] = 0
            }
        }
    }
    return res
}

//另外一种解法，对数组排序后，先遍历x，再去遍历剩余区域的left和right，通过l和r向中间靠拢来求解
//Time complexity O(n^2)
//Space complexity O(1)
func threeSum2(nums []int) [][]int {
    res, target := [][]int{}, 0
    if len(nums) < 3 {
        return res
    }

    //将nums排序，方便后续去重
    sort.Ints(nums)

    for i, x := range nums[:len(nums)-2]{
        //判重，重复的x没有意义
        if i>0 && nums[i-1] == x {
            continue
        }
        left, right := i+1, len(nums)-1

        for left < right {
            sum := x + nums[left] + nums[right]
            if sum < target {
                left++
            } else if sum > target {
                right--
            } else {
                res = append(res, []int{x, nums[left], nums[right]})
                //下面两个for进行去重处理，注意不要越界
                for left < right && nums[left] == nums[left+1]{
                    left++
                }
                for left < right && nums[right] == nums[right-1]{
                    right--
                }
                left++
                right--
            }
        }
    }
    return res
}

func main(){
    nums := []int{-1, -2, 2, 0, -2, 1, 2, -1, -4, 2, 4, 4}
    //nums := []int{0,0,0,0,0,0,0,0,0,0,0,0}
    ret := threeSum1(nums)
    fmt.Println(ret)
    ret = threeSum2(nums)
    fmt.Println(ret)
}

