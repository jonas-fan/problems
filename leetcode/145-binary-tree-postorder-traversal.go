// #tree #binary-tree #depth-first-search

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    out := []int{}
    stack := []*TreeNode{}
    last := &TreeNode{}

    for root != nil || len(stack) > 0 {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            node := stack[len(stack)-1]

            if node.Right == nil || node.Right == last {
                stack = stack[:len(stack)-1]
                out = append(out, node.Val)
                last = node
            } else {
                root = node.Right
            }
        }
    }

    return out
}
