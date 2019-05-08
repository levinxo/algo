package main

import "fmt"
import "strings"
import "sort"

//有效的字母异位词
//leetcode: https://leetcode.com/problems/valid-anagram/

//先排序，再比较
//Time complexity O(n*logn)
//Space complexity O(n)
func isAnagram1(s string, t string) bool {

    lambda := func(x string) string {
	xArr := strings.Split(x, "")
	sort.Strings(xArr)
	return strings.Join(xArr, "")
    }

    return lambda(s) == lambda(t)
}

//使用切片，但其实是hashmap的思路
//Time complexity O(n)
//Space complexity O(1)
func isAnagram2(s string, t string) bool {

    if len(s) != len(t){
	return false
    }

    arr := make([]int, 26)

    for _, v := range s {
	arr[v-'a']++
    }

    for _, v := range t {
	if arr[v-'a'] < 0 {
	    return false
	} else {
	    arr[v-'a']--
	}
    }

    for _, v := range arr {
	if v != 0 {
	    return false
	}
    }
    return true
}

func main(){
    s := "anagram"
    t := "nagaram"
    ret := isAnagram2(s, t)
    fmt.Println(ret)
}

