package lowestCommonAncestorrr

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type value struct {
	pv     int
	qv     int
	pb     bool
	qb     bool
	answer *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	v := new(value)
	v.pv = p.Val
	v.qv = q.Val
	v.answer = new(TreeNode)
	v.answer.Val = -1

	v.Find(root.Left)
	if v.answer.Val == -1 && (v.pb || v.qb) {
		return root
	} else {
		v.Find(root.Right)
	}
	return v.answer
}

func (s *value) Find(root *TreeNode) {
	if root != nil {
		s.Find(root.Left)

		if root.Val == s.qv {
			s.qb = true

		}
		if root.Val == s.pv {
			s.pb = true
		}
		if (s.pb && s.qb) && s.answer.Val == -1 {
			s.answer = root
			return
		}
		s.Find(root.Right)
	}
	return
}
