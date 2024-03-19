package isEvenOddTree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LevelTree struct {
	level int
	max   int
	lMap  map[int][]int
}

func LevelTreeInit() LevelTree {
	l := new(LevelTree)
	l.level = 0
	l.max = 0
	l.lMap = make(map[int][]int)
	return *l
}

func levelOrder(root *TreeNode) [][]int {
	l := LevelTreeInit()

	result := make([][]int, 0)
	l.lorder(root)
	for i := 0; i < l.max; i++ {
		result = append(result, l.lMap[i])
	}

	return result
}

func (l *LevelTree) lorder(root *TreeNode) {
	if root != nil {
		l.lMap[l.level] = append(l.lMap[l.level], root.Val)
		l.level++
		if l.level > l.max {
			l.max = l.level
		}
		l.lorder(root.Left)
		l.lorder(root.Right)
		l.level--
	}
	return
}

func isEvenOddTree(root *TreeNode) bool {
	matrix := levelOrder(root)
	fmt.Println(matrix)
	for i := 0; i < len(matrix); i++ {
		if i%2 == 0 {
			// 偶数 奇数 递增
			if len(matrix[i]) == 1 {
				if matrix[i][0]%2 == 0 {
					return false
				}
			} else {
				for j := 0; j < len(matrix[i])-1; j++ {
					if matrix[i][j]%2 == 0 || matrix[i][j] >= matrix[i][j+1] || matrix[i][j+1]%2 == 0 {
						return false
					}
				}
			}
		} else {
			//奇数 偶数 递减
			if len(matrix[i]) == 1 {
				if matrix[i][0]%2 != 0 {
					return false
				}
			} else {
				for j := 0; j < len(matrix[i])-1; j++ {
					if matrix[i][j]%2 != 0 || matrix[i][j] <= matrix[i][j+1] || matrix[i][j+1]%2 != 0 {
						return false
					}
				}
			}
		}
	}
	return true
}
