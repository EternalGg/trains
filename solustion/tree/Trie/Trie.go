package Trie

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	t := new(Trie)
	t.children = [26]*Trie{}
	return *t
}

func (this *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		now := word[i] - 'a'
		if this.children[now] == nil {
			t := new(Trie)
			this.children[now] = t
		}
		this = this.children[now]
	}
	this.isEnd = true
}

func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if this.children[word[i]-'a'] != nil {
			this = this.children[word[i]-'a']
		} else {
			return false
		}
	}
	if this.isEnd == true {
		return true
	} else {
		return false
	}
}

func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if this.children[prefix[i]-'a'] != nil {
			this = this.children[prefix[i]-'a']
		} else {
			return false
		}
	}
	return true
}
