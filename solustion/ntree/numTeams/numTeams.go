package numTeams

type Node struct {
	Val int
	//layer    int
	Biger    bool
	Children []*Node
}

func numTeams(rating []int) int {
	//headList, length := make([]Node, 0), len(rating)
	//
	//for i := 0; i < length-2; i++ {
	//	n := Node{Val: rating[i]}
	//	headList = append(headList, n)
	//}
	//
	//for i, node := range headList {
	//	for j := i; j < length; j++ {
	//		n := Node{Val: rating[j]}
	//		// 小于上级
	//		if rating[j] < node.Val {
	//			n.Biger = false
	//		} else {
	//			n.Biger = true
	//		}
	//		node.Children = append(node.Children, &n)
	//	}
	//}
	//for i, node := range headList {
	//	for j := 0; j < len(node.Children); j++ {
	//
	//	}
	//}
	return 0
}

// 2,5,3,4,1
func NumTeams2(rating []int) int {
	result, length := 0, len(rating)
	for i := 1; i < length-1; i++ {
		leftSmaller, leftBigger := 0, 0
		for j := i - 1; j >= 0; j-- {
			if rating[j] > rating[i] {
				leftBigger++
			} else {
				leftSmaller++
			}
		}
		for j := i + 1; j < length; j++ {
			// 如果后面的比当前小
			if rating[j] < rating[i] {
				result += leftBigger
			}
			// 如果后面的比当前大
			if rating[j] > rating[i] {
				result += leftSmaller
			}
		}
	}
	return result
}
