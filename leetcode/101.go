// https://leetcode.com/problems/symmetric-tree/

func isSymmetric(node *TreeNode) bool {
	return mirrorCheck(node, node)
}

func mirrorCheck(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if (l == nil && r != nil) || (r == nil && l != nil) {
		return false
	}

	return l.Val == r.Val && mirrorCheck(l.Left, r.Right) && mirrorCheck(l.Right, r.Left)
}