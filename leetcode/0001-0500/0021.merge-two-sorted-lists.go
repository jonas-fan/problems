// #linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head := ListNode{}
    node := &head

    for list1 != nil || list2 != nil {
        if list1 == nil {
            node.Next = list2
            list2 = nil
        } else if list2 == nil {
            node.Next = list1
            list1 = nil
        } else if list1.Val < list2.Val {
            node.Next = list1
            list1 = list1.Next
        } else {
            node.Next = list2
            list2 = list2.Next
        }

        node = node.Next
    }

    return head.Next
}
