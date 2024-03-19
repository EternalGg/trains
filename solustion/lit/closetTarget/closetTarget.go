package closetTarget

type circle struct {
	val    string
	next   *circle
	before *circle
}

func closetTarget(words []string, target string, startIndex int) int {
	if target == words[startIndex] {
		return 0
	}
	head := new(circle)
	node := head
	last := head
	for i := 0; i < len(words)-1; i++ {
		node.val = words[i]
		node.before = last
		newOne := new(circle)
		node.next = newOne
		last = node
		node = node.next
	}
	node.val = words[len(words)-1]
	node.next = head
	node.before = last
	head.before = node

	open, key := new(circle), 0
	for key != startIndex {
		head = head.next
		key++
	}
	open = head
	openCopy := open

	pos, ne := 0, 0
	for pos != 2*len(words) {
		pos++
		openCopy = openCopy.next
		if openCopy.val == target {
			break
		}
	}
	openCopy1 := open

	for ne != 2*len(words) {
		ne++
		openCopy1 = openCopy1.before
		if openCopy1.val == target {
			break
		}
	}
	if pos == (2 * len(words)) {
		return -1
	}
	if pos > ne {
		return ne
	} else {
		return pos
	}
}
