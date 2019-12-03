package q720

type TrieNode struct {
	Children map[rune]*TrieNode
	End int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
		End:      0,
	}
}

type Trie struct {
	Root  *TrieNode
	Words []string
}

func (t *Trie)Insert(w string, index int) {
	node := t.Root
	chars := []rune(w)

	for i := 0; i < len(chars); i++ {
		if _, ok := node.Children[chars[i]]; !ok {
			node.Children[chars[i]] = NewTrieNode()
		}
		node = node.Children[chars[i]]
	}
	node.End = index
}

func (t *Trie)DFS() string{
	ans := ""
	stack := []*TrieNode{t.Root}

	for {
		if len(stack) < 1 {
			break
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.End == 0 && node != t.Root {
			continue
		}
		if node != t.Root {
			word := t.Words[node.End-1]
			// compare sort
			if len(word) > len(ans) || (len(word) == len(ans) && word < ans) {
				ans = word
			}
		}

		for _, v := range node.Children {
			stack = append(stack, v)
		}
	}

	return ans
}

func NewTrie() *Trie{
	return &Trie{
		Root:  NewTrieNode(),
		Words: nil,
	}
}

func LongestWord(words []string) string{
	trie := NewTrie()

	for i := 0; i < len(words); i++ {
		trie.Insert(words[i], i+1)
	}
	trie.Words = words
	return trie.DFS()
}

