package replaceWords

import "fmt"

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

func (this *Trie) Search(word string) (bool, int) {
	for i := 0; i < len(word); i++ {
		if this.children[word[i]-'a'] != nil {
			this = this.children[word[i]-'a']
			if this.isEnd {
				return true, i
			}
		} else {
			return false, -1
		}
	}
	if this.isEnd == true {
		return true, len(word) - 1
	} else {
		return false, -1
	}
}

func ReplaceWords(dictionary []string, sentence string) string {
	result := ""
	t := Constructor()
	for i := 0; i < len(dictionary); i++ {
		t.Insert(dictionary[i])
	}

	last := ""
	for i := 0; i < len(sentence); i++ {
		if sentence[i] != ' ' {
			last += string(sentence[i])
		} else {
			if now, isExit := t.Search(last); now == true {
				result += last[:isExit+1]
			} else {
				result += last
			}
			result += " "
			last = ""
		}
	}
	if now, isExit := t.Search(last); now == true {
		result += last[:isExit+1]
	} else {
		result += last
	}
	fmt.Println(result)
	return result
}
