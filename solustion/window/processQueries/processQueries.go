package processQueries

type ListNode struct {
	Val  int
	Next *ListNode
}

func ProcessQueries(queries []int, m int) []int {
	mList := new(ListNode)
	cNode := mList
	for i := 1; i < m; i++ {
		mList.Val = i
		node := new(ListNode)
		mList.Next = node
		mList = mList.Next
	}
	mList.Val = m

	result := make([]int, len(queries))
	for i, query := range queries {
		if query == cNode.Val {
			result[i] = 0
		} else {
			result[i] = searchAndCarry(cNode, query)
			node := new(ListNode)
			node.Val = query
			node.Next = cNode
			cNode = node
		}
	}
	return result
}

func searchAndCarry(node *ListNode, value int) int {
	postion := 1
	for node.Next != nil {
		if node.Next.Val == value {
			if node.Next.Next != nil {
				node.Next = node.Next.Next
			} else {
				node.Next = nil
			}
			return postion
		}
		node = node.Next
		postion++
	}
	return postion
}
