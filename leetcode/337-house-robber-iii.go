/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rob(root *TreeNode) int {
    return max(sum(root))
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func sum(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }

    lrobbed, lsafe := sum(root.Left)
    rrobbed, rsafe := sum(root.Right)
    robbed := root.Val + lsafe + rsafe
    safe := max(lrobbed, lsafe) + max(rrobbed, rsafe)

    return robbed, safe
}
