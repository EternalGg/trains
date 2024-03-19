package WordDictionary

type WordDictionary struct {
	StrMap map[int][]string
}

func Constructor() WordDictionary {
	return WordDictionary{
		StrMap: map[int][]string{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.StrMap[len(word)] = append(this.StrMap[len(word)], word)
}

func (this *WordDictionary) Search(word string) bool {
	for _, str := range this.StrMap[len(word)] {
		if StrCompare(str, word) {
			return true
		}
	}
	return false
}

func StrCompare(str1, str2 string) bool {
	for i := 0; i < len(str1); i++ {
		if str2[i] == '.' {
			continue
		}
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}
