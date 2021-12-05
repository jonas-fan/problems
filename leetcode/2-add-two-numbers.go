/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head *ListNode
    var node **ListNode = &head
    var carry int

    for l1 != nil || l2 != nil || carry > 0 {
        sum := listValue(l1) + listValue(l2) + carry

        carry = sum / 10
        *node = &ListNode{
            Val: sum % 10,
        }

        node = &(*node).Next
        l1 = listNext(l1)
        l2 = listNext(l2)
    }

    return head
}

func listValue(head *ListNode) int {
    if head == nil {
        return 0
    }

    return head.Val
}

func listNext(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    return head.Next
}
