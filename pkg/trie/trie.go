package trie

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

type TrieNode struct {
	Children map[rune]*TrieNode
	IsWord   bool
}

func NewTrieNode() *TrieNode {
	t := &TrieNode{}
	t.Children = map[rune]*TrieNode{}
	return t
}

func (t *Trie) Add(word string) {
	curr := t.root

	for _, char := range word {
		if curr.Children[char] == nil {
			curr.Children[char] = NewTrieNode()
		}
		curr = curr.Children[char]
	}

	curr.IsWord = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root

	for _, char := range word {
		if curr.Children[char] == nil {
			return false
		}
		curr = curr.Children[char]
	}

	return curr.IsWord
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root

	for _, char := range prefix {
		if curr.Children[char] == nil {
			return false
		}
		curr = curr.Children[char]
	}

	return true
}

func (t *Trie) Remove(word string) {
	removeHelper(t.root, word, 0)
}

func removeHelper(node *TrieNode, word string, idx int) bool {
	if node == nil {
		return false
	}

	if idx == len(word) {
		node.IsWord = false
	} else {
		char := rune(word[idx])
		isLeaf := removeHelper(node.Children[char], word, idx+1)

		// If the node is a leaf, then it can be safely deleted.
		if isLeaf {
			delete(node.Children, char)
		}
	}

	return len(node.Children) == 0
}
