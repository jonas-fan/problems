/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getDecimalValue(head *ListNode) int {
    out := 0

    for ; head != nil; head = head.Next {
        if out != 0 {
            out = out<<1 + head.Val
        } else {
            out = head.Val
        }
    }

    return out
}
