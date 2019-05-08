package main

import "fmt"

//滑动窗口最大值
//leetcode: https://leetcode.com/problems/sliding-window-maximum/

//Use Deque (双端队列)
//Time complexity O(n)
//Space complexity O(k)
func maxSlidingWindow(nums []int, k int) []int {
    var res, window []int

    for index := 0; index < len(nums); index++ {
        //窗口滑动后，若超过k个，就要从左侧移除掉多余的窗口元素
	if index >= k && window[0] <= index - k {
	    window = window[1:]
	}
        //将当前循环到的index对应value和window的最后一位索引的对应value相比，
        //若大于它，就把window最后一位删掉，因为之后该位对应的value再也没有机会是最大的了
	for len(window) > 0 && nums[window[len(window) - 1]] < nums[index] {
	    window = window[:len(window) - 1]
	}
        //将当前index加入到window中
	window = append(window, index)
        //生成res数组
	if index >= k - 1 {
	    res = append(res, nums[window[0]])
	}
    }
    return res
}

func main(){
    arr := []int{1,3,-1,-3,5,3,6,7}
    k := 3
    res := maxSlidingWindow(arr, k)
    fmt.Println(res)
}

