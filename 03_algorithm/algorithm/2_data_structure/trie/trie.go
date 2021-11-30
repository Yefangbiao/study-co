package trie

type Trie struct {
	isEnd bool
	child map[byte]*Trie
}

func (t *Trie) Insert(str string) {
	cur := t
	for i := 0; i < len(str); i++ {
		if cur.child[str[i]] == nil {
			cur.child[str[i]] = &Trie{
				isEnd: false,
				child: map[byte]*Trie{},
			}
		}
		cur = cur.child[str[i]]
	}
	cur.isEnd = true
}

func (t *Trie) Exist(str string) bool {
	cur := t
	for i := 0; i < len(str); i++ {
		if cur.child[str[i]] == nil {
			return false
		}
		cur = cur.child[str[i]]
	}
	return cur.isEnd
}
