/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func next(node *ListNode) *ListNode {
    if node == nil {
        return nil
    }

    return node.Next
}

func eval(node *ListNode) int {
    if node == nil {
        return 0
    }

    return node.Val
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    node := &ListNode{}
    head := node
    digit := 0

    for l1 != nil || l2 != nil || digit != 0 {
        digit += eval(l1) + eval(l2)

        node.Next = &ListNode{
            Val:  digit % 10,
        }

        node = node.Next
        digit /= 10
        l1 = next(l1)
        l2 = next(l2)
    }

    return head.Next
}
