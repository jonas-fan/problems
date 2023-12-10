// #linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    head = &ListNode{Next: head}

    first := head
    last := head
    node := head.Next

    for node != nil {
        // take the last N+1 nodes
        last.Next = node
        last = last.Next

        if n > 0 {
            n--
        } else {
            first = first.Next
        }

        node = node.Next
    }

    // remove the N-th node
    first.Next = first.Next.Next

    return head.Next
}
