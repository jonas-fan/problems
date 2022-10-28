/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode, out []int) []int {
    if root == nil {
        return out
    }

    out = dfs(root.Left, out)
    out = append(out, root.Val)
    out = dfs(root.Right, out)

    return out
}

func inorderTraversal(root *TreeNode) []int {
    return dfs(root, []int{})
}
