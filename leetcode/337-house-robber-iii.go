/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var cache = map[*TreeNode]int{}

func rob(root *TreeNode) int {
    return max(sum(root, false), sum(root, true))
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func sum(root *TreeNode, robbed bool) int {
    if root == nil {
        return 0
    }

    if robbed {
        return root.Val + sum(root.Left, false) + sum(root.Right, false)
    }

    if val, ok := cache[root]; ok {
        return val
    }

    max := max(sum(root.Left, true), sum(root.Left, false)) +
        max(sum(root.Right, true), sum(root.Right, false))

    cache[root] = max

    return max
}
