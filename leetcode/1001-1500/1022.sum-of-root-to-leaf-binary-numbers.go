/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(node *TreeNode, base int) int {
    if node == nil {
        return 0
    }

    base = (base << 1) | node.Val

    // leaf?
    if node.Left == nil && node.Right == nil {
        return base
    }

    return dfs(node.Left, base) + dfs(node.Right, base)
}

func sumRootToLeaf(root *TreeNode) int {
    return dfs(root, 0)
}
