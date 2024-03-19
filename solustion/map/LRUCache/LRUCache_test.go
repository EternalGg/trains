package LRUCache

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Get(1)
	c.Put(3, 3)
	c.Get(2)
}
