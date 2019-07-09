package main

import (
    "fmt"
    "strconv"
)

//字符串相加
//leetcode: https://leetcode.com/problems/add-strings/
func addStrings(num1 string, num2 string) string {

    length1 := len(num1)
    length2 := len(num2)

    carry := 0
    data := ""

    for length1 > 0 || length2 > 0 || carry > 0 {
        plus1, plus2 := 0, 0

        //取最后一位数字进行相加
        if length1 > 0 {
            plus1, _ = strconv.Atoi(string(num1[length1-1]))
            length1--
            num1 = num1[:length1]
        }
        if length2 > 0 {
            plus2, _ = strconv.Atoi(string(num2[length2-1]))
            length2--
            num2 = num2[:length2]
        }

        //加之后利用取模和整除法来获取最后一位和进位
        plus := plus1 + plus2 + carry
        data = strconv.Itoa(plus%10) + data
        carry = plus/10
    }
    return data
}

func main(){
    num1 := "1687365188"
    num2 := "48746289"
    ret := addStrings(num1, num2)
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(ret)

    num1 = "2687365188"
    num2 = "9848746289"
    ret = addStrings(num1, num2)
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(ret)

    num1 = "87365188"
    num2 = "9848746289"
    ret = addStrings(num1, num2)
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(ret)

    num1 = "1111111111"
    num2 = "3333333323"
    ret = addStrings(num1, num2)
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(ret)
}

