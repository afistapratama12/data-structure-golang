package trie

import "fmt"

// trie implementation

/* example:
					root
		         /    |   \
				j     s    t
			   /      |     \
			  o       a      i
			 /|       |       \
            h s       m        m
            | |
			n e
data : jo, joh, john, jose, sam, tim

- Each node has ALBHABET_SIZE=26 children. Each child node is a trie node itself and has ALBHABET_SIZE=26 children.
- Each node is at an index in its parent’ children node array and represents an ASCII character. For eg for a particular node, the first non-nil children node will mean the presence of char ‘a’,  second non-nil children node means the presence of ‘b’ and so on. Absence of a child at an index means no value
- Each node also has a boolean field indicating whether the node is the end of word or not
- The root node is the starting node and has ALBHABET_SIZE=26 children. root is associated with an empty value
*/

const (
	ALPHABET_SIZE = 26
)

type TrieNode struct {
	childrens [ALPHABET_SIZE]*TrieNode
	isWordEnd bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &TrieNode{}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
}

func (t *Trie) find(word string) bool {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	if current.isWordEnd {
		return true
	}
	return false
}

func initTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func trie() {
	trie := initTrie()

	words := []string{"sam", "john", "tim", "jose", "rose", "cat", "dog", "dogg", "roses"}

	for i := 0; i < len(words); i++ {
		trie.insert(words[i])
	}

	fmt.Println(trie)

	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	for i := 0; i < len(wordsToFind); i++ {
		found := trie.find(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}
}
