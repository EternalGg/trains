package constructMaximumBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type cmbt struct {
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	result := new(TreeNode)

	return result
}

func createMaxBinaryTree(root *TreeNode, nums []int) {
	//max := findMax(nums)
	//left := nums[]
	return
}

func findMax(nums []int) int {
	max, maxKey := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxKey = i
		}
	}
	return maxKey
}
