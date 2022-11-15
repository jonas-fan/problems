// #linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
    node := head
    prev := &ListNode{Next: node}

    head = prev

    for node != nil {
        if node.Val == val {
            prev.Next = nil
        } else {
            prev.Next = node
            prev = prev.Next
        }

        node = node.Next
    }

    return head.Next
}
