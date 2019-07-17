package main

import "fmt"

//单词搜索 II
//leetcode: https://leetcode.com/problems/word-search-ii/
//使用trie树和dfs
//代码还可以精简一下：把trie树的遍历流程也放到递归里去

type Trie struct {
    children map[rune]*Trie     //或者用数组也行
    isWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{map[rune]*Trie{}, false}
}

/** Inserts a word into the trie. */
//Time complexity O(n)
//Space complexity O(n)
func (this *Trie) Insert(word string)  {
    node := this
    for _, char := range word {
        nextNode, ok := node.children[char]
        if !ok {
            obj := Constructor()
            nextNode = &obj
            node.children[char] = nextNode
        }
        node = nextNode
    }
    node.isWord = true
}

/** Returns if the word is in the trie. */
//Time complexity O(n)
//Space complexity O(1)
func (this *Trie) Search(word string) bool {
    node := this
    for _, char := range word {
        nextNode, ok := node.children[char]
        if !ok {
            return false
        }
        node = nextNode
    }
    return node.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
//Time complexity O(n)
//Space complexity O(1)
func (this *Trie) StartsWith(prefix string) bool {
    node := this
    for _, char := range prefix {
        nextNode, ok := node.children[char]
        if !ok {
            return false
        }
        node = nextNode
    }
    return true
}

var res map[string]bool

func findWords(board [][]byte, words []string) []string {
    res = map[string]bool{}

    trie := Constructor()
    for _, str := range words {
        trie.Insert(str)
    }

    m := len(board)
    n := len(board[0])

    var visited [][]bool
    for i:=0; i<m; i++ {
        list := []bool{}
        for j:=0; j<n; j++ {
            list = append(list, false)
        }
        visited = append(visited, list)
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            dfs(board, visited, "", i, j, trie)
        }
    }

    tmp := []string{}

    for word, _ := range res {
        tmp = append(tmp, word)
    }
    return tmp
}

func dfs(board [][]byte, visited [][]bool, word string, i int, j int, trie Trie){
    if i<0 || i>=len(board) || j<0 || j>=len(board[0]) || visited[i][j] == true {
        return
    }

    word = word + string(board[i][j])
    if !trie.StartsWith(word) {
        return
    }

    if trie.Search(word) {
        res[word] = true
    }

    visited[i][j] = true
    dfs(board, visited, word, i-1, j, trie)
    dfs(board, visited, word, i+1, j, trie)
    dfs(board, visited, word, i, j-1, trie)
    dfs(board, visited, word, i, j+1, trie)
    visited[i][j] = false
}

func main(){

    words := []string{"oath", "pea", "eat", "rain"}
    board := [][]byte{
        []byte{'o', 'a', 'a', 'n'},
        []byte{'e', 't', 'a', 'e'},
        []byte{'i', 'h', 'k', 'r'},
        []byte{'i', 'f', 'l', 'v'},
    }

    //words := []string{"baa", "abba", "baab", "aba"}
    //board := [][]byte{
    //    []byte{'b'},
    //    []byte{'a'},
    //    []byte{'b'},
    //    []byte{'b'},
    //    []byte{'a'},
    //}

    //words := []string{"a"}
    //board := [][]byte{
    //    []byte{'a'},
    //    []byte{'a'},
    //}

    ret := findWords(board, words)
    fmt.Println(ret)
}

