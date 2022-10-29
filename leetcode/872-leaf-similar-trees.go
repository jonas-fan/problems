/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func equals(lhs []int, rhs []int) bool {
    if len(lhs) != len(rhs) {
        return false
    }

    for i := range lhs {
        if lhs[i] != rhs[i] {
            return false
        }
    }

    return true
}

func dfs(node *TreeNode, out []int) []int {
    if node == nil {
        return out
    }

    // leaf?
    if node.Left == nil && node.Right == nil {
        return append(out, node.Val)
    }

    out = dfs(node.Left, out)
    out = dfs(node.Right, out)

    return out
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    return equals(dfs(root1, []int{}), dfs(root2, []int{}))
}
