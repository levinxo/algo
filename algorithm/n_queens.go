package main

import "fmt"
import "strings"

//N皇后
//leetcode: https://leetcode.com/problems/n-queens/

//利用递归+剪枝
//关键在于剪枝，一个皇后占位后，它的行和列、斜线、反斜线都无法再被占用
//使用一个map来存储被占用的点规则信息
//Time complexity O(n^2)，存疑
//Space complexity O(n^2)

type DataSet struct {
    Col map[int] bool
    Slash map[int] bool
    BSlash map[int] bool
}

var dataSet DataSet
var allQueenString [][]string

func solveNQueens(n int) [][]string{
    if n < 1 {
        return [][]string{}
    }

    dataSet = DataSet{}
    dataSet.Col = make(map[int] bool)
    dataSet.Slash = make(map[int] bool)
    dataSet.BSlash = make(map[int] bool)
    allQueenString = [][]string{}

    helper(0, n, []int{})
    return allQueenString
}

func helper(row int, n int, colPosition []int) {
    if row == n {
        allQueenString = append(allQueenString, printStr(n, colPosition))
        return
    }

    for col:=0; col<n; col++ {
        _, colExist := dataSet.Col[col]
        _, slashExist := dataSet.Slash[(row+col)]
        _,  bSlashExist := dataSet.BSlash[(row-col)]
        if colExist || slashExist || bSlashExist {
            continue
        }

        dataSet.Col[col] = true
        dataSet.Slash[(row+col)] = true
        dataSet.BSlash[(row-col)] = true

        helper(row+1, n, append(colPosition, col))

        delete(dataSet.Col, col)
        delete(dataSet.Slash, row+col)
        delete(dataSet.BSlash, row-col)
    }
}

func printStr(n int, pos []int) []string {
    arr := []string{}
    for j:=0; j<n; j++ {
        rowStr := strings.Repeat(".", pos[j]) + "Q" + strings.Repeat(".", n-1-pos[j])
        arr = append(arr, rowStr)
    }
    return arr
}

func main(){
    n := 7
    ret := solveNQueens(n)
    fmt.Println("total result:", len(ret))
    fmt.Println("output:")
    fmt.Println(ret)
}

