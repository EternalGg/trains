package deleteNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil { //没找到，返回nil
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key) //向左子树搜索
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key) //向右子树搜索
	} else { //找到了需要删除的节点
		if root.Left == nil { //第一种情况，被删除节点只有右子树，直接返回左子树
			return root.Right
		}
		if root.Right == nil { //第二种情况，被删除节点只有左子树，直接返回右子树
			return root.Left
		}

		max := root.Left
		for max.Right != nil {
			max = max.Right
		}
		max.Right = root.Right
		root = root.Left
	}
	return root
}
