package main

import "fmt"

//解数独
//leetcode: https://leetcode.com/problems/sudoku-solver/

//利用递归
//@todo 剪枝
//剪枝方法：
//1.递归时，记住前一次递归的i和j，下次递归就不要重新开始了
//2.递归开始之前，将数独开始遍历的点排一个优先级，可选数格子少的优先级更高
//3.参考一下dancing link数据结构
//Time complexity O(?)
//Space complexity O(?)

func solveSudoku(board [][]byte) {
    if board == nil || len(board) < 1 {
        return
    }
    ret := helper(&board, 0)
    fmt.Println(ret)
    printStr(board)
}

func helper(board *[][]byte, row int) bool {
    for i:=row; i<len(*board); i++ {
        for j:=0; j<len((*board)[0]); j++ {
            if (*board)[i][j] == '.' {
                for c:=byte('1'); c<=byte('9'); c++ {
                    if isValid(board, i, j, c) {
                        (*board)[i][j] = c
                        if helper(board, i) {
                            return true
                        } else {
                            //go back
                            (*board)[i][j] = '.'
                        }
                    }
                }
                return false
            }
        }
    }
    return true
}

func isValid(board *[][]byte, i, j int, c byte) bool {
    for x:=0; x<9; x++ {
        if (*board)[i][x] == c {return false}
        if (*board)[x][j] == c {return false}
        //通过3*(i/3)获得3x3格子的坐标原点
        //再通过x/3和x%3来逐次遍历这9个格子
        if (*board)[3*(i/3)+x/3][3*(j/3)+x%3] == c {return false}
    }
    return true
}

func printStr(arr [][]byte){
    for _, row := range arr {
        for _, val := range row {
            fmt.Printf("%s ", string(val))
        }
        fmt.Printf("\n")
    }
}

func main(){
    arr := [][]byte{
        []byte{'5','3','.','.','7','.','.','.','.'},
        []byte{'6','.','.','1','9','5','.','.','.'},
        []byte{'.','9','8','.','.','.','.','6','.'},
        []byte{'8','.','.','.','6','.','.','.','3'},
        []byte{'4','.','.','8','.','3','.','.','1'},
        []byte{'7','.','.','.','2','.','.','.','6'},
        []byte{'.','6','.','.','.','.','2','8','.'},
        []byte{'.','.','.','4','1','9','.','.','5'},
        []byte{'.','.','.','.','8','.','.','7','9'},
    }
    solveSudoku(arr)
}

