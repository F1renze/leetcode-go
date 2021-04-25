package q208

import (
	"testing"

	"github.com/f1renze/leetcode-go/dskit"
)

func TestTrie_StartsWith(t *testing.T) {
	trie := Constructor()

	trie.Insert("hello")
	dskit.Equals(t, false, trie.Search("hell"))
	dskit.Equals(t, false, trie.StartsWith("helloa"))
}
