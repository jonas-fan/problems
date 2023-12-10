// #linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var node = head
    var prev *ListNode = nil

    for node != nil  {
        next := node.Next
        node.Next = prev
        prev = node
        node = next
    }

    return prev
}
