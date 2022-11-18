// #tree #binary-tree #depth-first-search

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    out := []int{}
    stack := []*TreeNode{}

    for root != nil || len(stack) > 0 {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            node := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            out = append(out, node.Val)

            if node.Right != nil {
                root = node.Right
            }
        }
    }

    return out
}
