// #linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
    mid := head
    node := head

    for node != nil {
        node = node.Next

        if node == nil {
            break
        }

        node = node.Next
        mid = mid.Next
    }

    return mid
}
