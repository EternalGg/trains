package MyHashMap

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type MyHashMap struct {
	Head   *Node
	MinKey int
	MaxKey int
	length uint8
}

func Constructor() MyHashMap {
	Map := new(MyHashMap)
	Map.Head = new(Node)
	Map.length = 0
	return *Map
}

func (this *MyHashMap) Put(key int, value int) {
	if this.length == 0 {
		this.Head.Key = key
		this.Head.Value = value
		this.MinKey = key
		this.MaxKey = key
		this.length++
	} else {
		switch {
		case this.MinKey == key:
			this.Head.Value = value
		case this.MinKey > key:
			newNode := new(Node)
			newNode.Key = key
			newNode.Value = value

			newNode.Next = this.Head
			this.Head = newNode
			this.MinKey = key
			this.length++
		case this.MaxKey == key:
			for this.Head.Next != nil {
				this.Head = this.Head.Next
			}
			this.Head.Value = value
		case this.MaxKey < key:
			newNode := new(Node)
			newNode.Key = key
			newNode.Value = value
			this.Head.Next = newNode
			this.MaxKey = key
			this.length++
		case this.MaxKey > key && this.MinKey < key:
			last := new(Node)
			for !(this.Head.Key <= key) {
				last = this.Head
				this.Head = this.Head.Next
			}
			if this.Head.Key == key {
				this.Head.Value = value
			} else {
				newNode := new(Node)
				newNode.Key = key
				newNode.Value = value
				last.Next = newNode
				newNode.Next = this.Head
				this.length++
			}
		}
	}
}

func (this *MyHashMap) Get(key int) int {
	if this.length == 0 {
		return -1
	}

	if key > this.MaxKey || key < this.MinKey {
		return -1
	}
	switch key {
	case this.MinKey:
		return this.Head.Value
	case this.MaxKey:
		for this.Head.Next != nil {
			this.Head = this.Head.Next
		}
		return this.Head.Value
	default:
		for this.Head != nil {
			if this.Head.Key == key {
				return this.Head.Value
			}
			this.Head = this.Head.Next
		}
		return this.Head.Value
	}

}

func (this *MyHashMap) Remove(key int) {
	if this.length == 0 {
		return
	}
	if this.length == 1 {
		this.Head = new(Node)
		this.length = 0
		this.MinKey = 0
		this.MaxKey = 0
		return
	}
	if key > this.MaxKey || key < this.MinKey {
		return
	}

	switch key {
	case this.MinKey:
		this.Head = this.Head.Next
		this.MinKey = this.Head.Key
	case this.MaxKey:
		last := new(Node)
		for this.Head.Next != nil {
			last = this.Head
			this.Head = this.Head.Next
		}
		last.Next = nil
		this.MaxKey = last.Key
	default:
		for this.Head != nil {
			if this.Head.Next.Key == key {
				Next := this.Head.Next.Next
				this.Head.Next = Next
				break
			}
		}
	}
	this.length--
}
