package main

import "fmt"

//比较含退格的字符串
//leetcode: https://leetcode.com/problems/backspace-string-compare/

//Use Stack
//Time complexity O(n)
//Space complexity O(n)
func backspaceCompare(S string, T string) bool {
    return buildStack(S) == buildStack(T)
}

func buildStack(s string) string {
    var a []byte

    for i:=0; i<len(s); i++ {
        if s[i] != '#' {
            a = append(a, s[i])
        } else if len(a) > 0 {
            a = a[0:len(a)-1]
        }
    }
    return string(a[:])
}

func main(){
    a := "ab#c"
    b := "ad#c"

    r := backspaceCompare(a, b)
    if r {
        fmt.Println("a==b")
    } else {
        fmt.Println("a!=b")
    }
}

