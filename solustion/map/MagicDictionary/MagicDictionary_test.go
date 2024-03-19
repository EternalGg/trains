package MagicDictionary

import "testing"

func TestConstructor(t *testing.T) {
	m := Constructor()
	m.BuildDict([]string{"hello", "leetcode"})
	m.Search("hhllo")
}
