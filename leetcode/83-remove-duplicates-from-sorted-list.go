// #linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    node := head
    prev := head

    node = node.Next

    for node != nil {
        if node.Val == prev.Val {
            prev.Next = node.Next
        } else {
            prev.Next = node
            prev = prev.Next
        }

        node = node.Next
    }

    return head
}
