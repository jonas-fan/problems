// #linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    fast := head
    slow := head
    cycled := false

    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if fast == slow {
            cycled = true
            break
        }
    }

    if !cycled {
        return nil
    }

    for head != slow {
        head = head.Next
        slow = slow.Next
    }

    return head
}
