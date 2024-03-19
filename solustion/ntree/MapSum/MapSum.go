package MapSum

import "fmt"

type MapSum struct {
	children [26]*MapSum
	isEnd    bool
	score    int
	val      int
}

func Constructor() MapSum {
	m := new(MapSum)
	m.children = [26]*MapSum{}
	return *m
}

func (this *MapSum) Search(word string) (bool, int) {
	for i := 0; i < len(word); i++ {
		if this.children[word[i]-'a'] != nil {
			this = this.children[word[i]-'a']
		} else {
			return false, 0
		}
	}
	if this.isEnd == true {
		return true, this.score
	} else {
		return false, 0
	}
}

func (this *MapSum) Insert(key string, val int) {
	if isExit, score := this.Search(key); isExit == true {
		for i := 0; i < len(key); i++ {
			now := key[i] - 'a'
			this.score -= score
			this.score += val
			fmt.Println(key[i], this.score)
			this = this.children[now]

		}
		this.score -= score
		this.score += val
		this.val = val
	} else {
		for i := 0; i < len(key); i++ {
			now := key[i] - 'a'
			if this.children[now] == nil {
				t := new(MapSum)
				this.children[now] = t
			}
			this.score += val
			this = this.children[now]
		}
		this.isEnd = true
		this.val = val
		this.score += val
	}
}

func (this *MapSum) Sum(prefix string) int {
	for i := 0; i < len(prefix); i++ {
		if this.children[prefix[i]-'a'] != nil {
			this = this.children[prefix[i]-'a']
		} else {
			return 0
		}
	}
	return this.score
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
