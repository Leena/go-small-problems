// https://leetcode.com/problems/add-two-numbers/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	current := sum
	total := 0

	for l1 != nil || l2 != nil || total != 0 {
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: total % 10}
		total /= 10
		current = current.Next
	}

	return sum.Next
}