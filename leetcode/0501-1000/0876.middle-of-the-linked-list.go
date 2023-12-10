// #linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
    fast := head
    slow := head

    for fast != nil {
        fast = fast.Next

        if fast == nil {
            break
        }

        fast = fast.Next
        slow = slow.Next
    }

    return slow
}
