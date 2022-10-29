/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(node *TreeNode, out []int) []int {
    if node == nil {
        return out
    }

    out = dfs(node.Left, out)
    out = dfs(node.Right, out)
    out = append(out, node.Val)

    return out
}

func postorderTraversal(root *TreeNode) []int {
    return dfs(root, []int{})
}
