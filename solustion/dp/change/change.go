package change

import (
	"sort"
)

//func Change(amount int, coins []int) int {
//	coinMap, count := make(map[int]bool), 0
//	newList := []int{}
//	for i := 0; i < len(coins); i++ {
//		if coins[i] == amount {
//			count++
//		}
//		if coins[i] < amount {
//			coinMap[coins[i]] = true
//			newList = append(newList, coins[i])
//		}
//	}
//	for len(coinMap) != 0 {
//		for key, _ := range coinMap {
//			for _, value := range newList {
//				if value+key < amount {
//					coinMap[value+key] = true
//				}
//				if value+key == amount {
//					count++
//				}
//			}
//			delete(coinMap, key)
//		}
//	}
//	return count
//}

var Count int

type Node struct {
	Val      int
	Count    int
	Children []*Node
}

func Change(amount int, coins []int) int {
	Count = 0
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})
	coinMap, remainMap := make(map[int]bool), map[int]bool{}
	for i := 0; i < len(coins); i++ {
		if coins[i] > amount {
			break
		}
		if coins[i] == amount {
			Count++
		}
		if coins[i] < amount {
			coinMap[coins[i]] = true
			remainMap[coins[i]] = true
		}
	}
	if len(remainMap) == 0 {
		return Count
	}

	headList := []Node{}
	for k, _ := range remainMap {
		CL := make([]*Node, 0)
		newNode := Node{Val: k, Count: k, Children: CL}
		headList = append(headList, newNode)
	}
	for _, Nodes := range headList {
		Insert(&Nodes, remainMap, amount)
		//fmt.Println(Nodes.Val)
	}
	return Count
}

func Insert(n *Node, m map[int]bool, amount int) {
	for k, _ := range m {
		if n.Val > k {
			continue
		}
		if k+n.Count < amount {
			CL := make([]*Node, 0)
			newNode := Node{Val: k, Count: k + n.Count, Children: CL}
			n.Children = append(n.Children, &newNode)
			Insert(&newNode, m, amount)
		}
		if k+n.Count >= amount {
			if k+n.Count == amount {
				//fmt.Println(k, n.Val, n.Count)
				Count++
			}
			return
		}
	}
}

func Change1(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	//外层循环是遍历每个面值的硬币
	//我们的思路是针对amount，计算出必须包含每个硬币的组合，然后累加
	//比如，先计算amount必须包含面值1的硬币，在计算必须包含面值为2的组合
	//是不是会重复呢，即包含面值1的组合中已经包括了包含面值2的，答案是不会重复，为何呢？
	//我们外层循环是从coin开始遍历的，比如amount=5，coins[5，2，1]先找到必须包含面值5的
	//这时候，会把dp[5]累加一个dp[0]，其他不会有影响，接下来，找必须包含面值2的

	//如果把两个for循环颠倒，则计算的是排列数，而不是组合数，即，2，1，1，1与1，1，1，2视为两种
	//这个要注意
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
