package main

import "fmt"

//有效的括号
//leetcode: https://leetcode.com/problems/valid-parentheses/

//Use Stack
//Time complexity O(n)
//Space complexity O(n)
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	sign := make(map[byte]byte)
	sign[')'] = '('
	sign['}'] = '{'
	sign[']'] = '['

	var stack []byte

	for i := 0; i < len(s); i++ {
		val := s[i]

		if _, ok := sign[val]; ok {
			// 如果是类似)}]这样的反括号的话，开始尝试消栈
			var last byte
			if len(stack) > 0 {
				last = stack[len(stack)-1:][0]
			}
			if last == sign[val] {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		} else {
			// 如果是类似({[这样的正括号的话，直接入栈
			stack = append(stack, val)
		}
	}

	return len(stack) == 0
}

func main() {
	//a := "()[]{}"
	a := "(])"

	r := isValid(a)
	if r {
		fmt.Println("valid")
	} else {
		fmt.Println("not valid")
	}
}
