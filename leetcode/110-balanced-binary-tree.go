/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func abs(num int) int {
    if num < 0 {
        return -num
    }

    return num
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func walk(node *TreeNode) (int, bool) {
    if node == nil {
        return 0, true
    }

    ldepth, lbalanced := walk(node.Left)
    rdepth, rbalanced := walk(node.Right)

    depth := 1 + max(ldepth, rdepth)
    balanced := lbalanced && rbalanced && (abs(ldepth-rdepth) < 2)

    return depth, balanced
}

func isBalanced(root *TreeNode) bool {
    _, balanced := walk(root)

    return balanced
}
