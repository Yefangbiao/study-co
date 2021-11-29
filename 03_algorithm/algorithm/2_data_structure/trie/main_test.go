package trie

import "testing"

func TestTrie_Exist(t *testing.T) {
	trie := &Trie{
		isEnd: false,
		child: map[byte]*Trie{},
	}

	trie.Insert("abc")
	trie.Insert("ab")
	trie.Insert("cc")
	trie.Insert("dd")

	if !trie.Exist("abc") {
		t.Errorf("abc is not exist")
	}
	if trie.Exist("ee") {
		t.Errorf("ee is not exist")
	}
}
