// #tree #binary-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
    out := []int{}
    stack := []*TreeNode{root}

    for len(stack) > 0 {
        node := stack[len(stack)-1]

        stack = stack[:len(stack)-1]

        if node != nil {
            out = append(out, node.Val)
            stack = append(stack, node.Right)
            stack = append(stack, node.Left)
        }
    }

    return out
}
