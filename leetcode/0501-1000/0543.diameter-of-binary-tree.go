/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func walk(node *TreeNode) (int, int) {
    if node == nil {
        return 0, 0
    }

    ldepth, ldia := walk(node.Left)
    rdepth, rdia := walk(node.Right)

    depth := 1 + max(ldepth, rdepth)
    dia := max(ldepth+rdepth, max(ldia, rdia))

    return depth, dia
}

func diameterOfBinaryTree(root *TreeNode) int {
    _, dia := walk(root)

    return dia
}
