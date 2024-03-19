package finalValueAfterOperations

import "testing"

func finalValueAfterOperations(operations []string) int {
	qa := 0
	for _, operation := range operations {
		switch operation {
		case "x++":
			qa++
		case "++x":
			qa++
		case "x--":
			qa--
		case "--x":
			qa--
		}
	}
	return qa
}
func TestFinalValueAfterOperations(t *testing.T) {

}
