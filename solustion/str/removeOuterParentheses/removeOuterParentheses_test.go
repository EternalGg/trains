package removeOuterParentheses

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
	RemoveOuterParentheses("(()())(())(()(()))")
}
