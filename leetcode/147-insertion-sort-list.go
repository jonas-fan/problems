/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertionSortList(head *ListNode) *ListNode {
    node := head
    head = &ListNode{}

    for node != nil {
        next := node.Next
        prev := head

        for prev.Next != nil && prev.Next.Val < node.Val {
            prev = prev.Next
        }

        node.Next = prev.Next
        prev.Next = node
        node = next
    }

    return head.Next
}
