package main

import (
    "fmt"
    "strings"
)

//151. 翻转字符串里的单词
//https://leetcode.com/problems/reverse-words-in-a-string/

//这个方法在go里面貌似性能不太好，可以考虑使用strings.Split和Join
//use stack
//Time Complexity O(k)
//Auxiliary Space Complexity O(k)
func reverseWords(s string) string {
    s = strings.Trim(s, " ")
    if s == "" {
        return s
    }

    var stack []string
    tmp, res := "", ""

    for _, c := range s {
        if c == ' ' {
            if tmp != "" {
                stack = append(stack, tmp)
                tmp = ""
            }
            continue
        }
        tmp += string(c)
    }

    //注意，最后一个单词也要存入stack
    stack = append(stack, tmp)

    for len(stack) > 0 {
        res += stack[len(stack)-1] + " "
        stack = stack[:len(stack)-1]
    }
    return strings.Trim(res, " ")
}

func main(){
    s := "  hello  world!  "
    s = reverseWords(s)

    fmt.Println(s)
}
