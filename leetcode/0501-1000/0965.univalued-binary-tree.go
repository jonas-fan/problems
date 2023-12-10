/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(node *TreeNode, val int) bool {
    if node == nil {
        return true
    }

    if node.Val != val {
        return false
    }

    return dfs(node.Left, val) && dfs(node.Right, val)
}

func isUnivalTree(root *TreeNode) bool {
    return dfs(root, root.Val)
}
