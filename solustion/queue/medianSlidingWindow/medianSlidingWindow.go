package medianSlidingWindow

import "fmt"

type doubleNode struct {
	Val  int
	Up   *doubleNode
	Down *doubleNode
}

type queueHaveMid struct {
	List    []doubleNode
	MidNode *doubleNode
}

func (q queueHaveMid) init(nums []int, k int) {
	queue := new(queueHaveMid)
	queue.List = make([]doubleNode, k)
}

func (q doubleNode) insert(num int) {

}

func (q doubleNode) delete(num int) {

}

func MedianSlidingWindow(nums []int, k int) []float64 {
	result, left, right := make([]float64, k), 0, k-1
	fmt.Println(left, right)
	return result
}
