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
    return max(sum(root, 0, false), sum(root, 0, true))
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func sum(root *TreeNode, val int, can bool) int {
    if root == nil {
        return val
    }

    if can {
        return root.Val + sum(root.Left, val, !can) + sum(root.Right, val, !can)
    }

    if val, ok := cache[root]; ok {
        return val
    }

    max := max(sum(root.Left, val, can), sum(root.Left, val, !can)) +
        max(sum(root.Right, val, can), sum(root.Right, val, !can))

    cache[root] = max

    return max
}
