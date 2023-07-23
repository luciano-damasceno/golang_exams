package main

import "fmt"

const AlphabeticLetters int = 26

// Node - Nodes of a tree holding the 26 alphabetic letters
type Node struct {
	children [AlphabeticLetters]*Node
	isEnd    bool
}

// Trie - Tree with a root node
type Trie struct {
	root *Node
}

type TrieInterface interface {
	addWord(string)
	searchWord(w string) bool
	removeWord(w string)
}

// Just initialize an empty trie with an empty root
func initTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (t *Trie) addWord(w string) {
	currentNode := t.root
	for i := 0; i < len(w); i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) searchWord(w string) bool {
	currentNode := t.root
	for i := 0; i < len(w); i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return true
}

func (n *Node) isItFreeNode() bool {
	for _, c := range n.children {
		if c != nil {
			return false
		}
	}
	return true
}

func (t *Trie) removeWord(word string) {
	if t.searchWord(word) {
		level := 0
		charIndex := word[level] - 'a'
		t.root.children[charIndex].removeWord(word, level)
	}
}

func (n *Node) removeWord(word string, level int) {
	// The reason why its -2 is because +1 is from word length starting at 1 and not 0 and +1 to get the children node.
	if level == len(word)-2 {
		childrenIndex := word[level+1] - 'a'
		n.children[childrenIndex].isEnd = false
		if n.children[childrenIndex].isItFreeNode() {
			n.children[childrenIndex] = nil
		}
	} else {
		nodeIndex := word[level] - 'a'
		fmt.Printf("%c\n", nodeIndex+'a')

		level++
		childrenIndex := word[level] - 'a'
		if n.children[childrenIndex] != nil {
			n.children[childrenIndex].removeWord(word, level)
			if n.children[childrenIndex].isItFreeNode() {
				n.children[childrenIndex] = nil
			}
		}
	}
}

// Needs fix in logics.
func printTrie(n Node) {
	for i, c := range n.children {
		if c != nil {
			if c.isEnd {
				fmt.Printf("%c\n", i+'a')
			} else {
				fmt.Printf("%c", i+'a')
			}
			printTrie(*c)
		}
	}
}

func main() {
	trie := initTrie()

	tt := []string{"lucianoo", "iris", "dunga", "gadelha", "lucian", "lucy"}

	for _, word := range tt {
		trie.addWord(word)
	}
	trie.removeWord("lucianoo")
	printTrie(*trie.root)
}
