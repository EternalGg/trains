package replaceWords

import "testing"

func TestConstructor(t *testing.T) {
	d := []string{"cat", "bat", "rat"}
	ReplaceWords(d, "the cattle was rattled by the battery")
}
