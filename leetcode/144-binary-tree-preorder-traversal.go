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

    out = append(out, node.Val)
    out = dfs(node.Left, out)
    out = dfs(node.Right, out)
    
    return out
}

func preorderTraversal(root *TreeNode) []int {
    return dfs(root, []int{})    
}
