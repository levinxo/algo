package main

import "fmt"

//实现 Trie (前缀树)
//leetcode: https://leetcode.com/problems/implement-trie-prefix-tree/

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

/**
 * Your Trie object will be instantiated and called as such:
 */
func main(){
    word := "google"
    prefix := "goo"

    obj := Constructor();
    obj.Insert(word);
    param_2 := obj.Search(word);
    param_3 := obj.StartsWith(prefix);

    if param_2 {
        fmt.Println(word)
    }
    if param_3 {
        fmt.Println(prefix)
    }
}

