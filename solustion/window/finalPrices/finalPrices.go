package finalPrices

type listNode struct {
	Val  int
	key  int
	Next *listNode
}

func FinalPrices(prices []int) []int {
	l := new(listNode)

	for i := 0; i < len(prices); i++ {
		node := l
		if node.Val == 0 {
			node.key = i
			node.Val = prices[i]
		} else {
			for node.Val >= prices[i] {
				prices[node.key] = node.Val - prices[i]
				if node.Next == nil {
					node.key = i
					node.Val = prices[i]
					break
				} else {
					node = node.Next
				}
			}
			newNode := new(listNode)
			newNode.key = i
			newNode.Val = prices[i]
			newNode.Next = node
			l = newNode
		}
	}
	return prices
}
