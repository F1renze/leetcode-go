package q208

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
		End:      false,
	}
}

type TrieNode struct {
	Children map[rune]*TrieNode
	End      bool
}

type Trie struct {
	Root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Root: NewTrieNode()}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	chars := []rune(word)
	node := t.Root

	for _, c := range chars {
		if _, ok := node.Children[c]; !ok {
			break
		}
		node = node.Children[c]
	}
	node.End = true
}

func (t *Trie) SearchPrefix(word string) *TrieNode {
	chars := []rune(word)
	node := t.Root

	for _, c := range chars {
		if _, ok := node.Children[c]; !ok {
			return nil
		}
		node = node.Children[c]
	}
	return node
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.End
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	node := t.SearchPrefix(prefix)
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
