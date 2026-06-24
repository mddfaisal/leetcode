package problem2

//
// https://leetcode.com/problems/add-two-numbers/
//

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	current := ret
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = int(sum / 10)
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	return ret.Next
}
