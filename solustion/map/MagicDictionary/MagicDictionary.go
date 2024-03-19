package MagicDictionary

type MagicDictionary struct {
	m map[int][]string
}

func Constructor() MagicDictionary {
	m := MagicDictionary{make(map[int][]string)}
	return m
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for i := 0; i < len(dictionary); i++ {
		this.m[len(dictionary[i])] = append(this.m[len(dictionary[i])], dictionary[i])
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	if len(this.m[len(searchWord)]) == 0 {
		return false
	} else {
		for _, value := range this.m[len(searchWord)] {
			diff := 0
			for i := 0; i < len(value); i++ {
				if value[i] != searchWord[i] {
					diff++
				}
				if diff > 1 {
					goto here
				}
			}
			if diff == 1 {
				return true
			}
		here:
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
