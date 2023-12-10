// #tree #binary-tree #binary-search-tree #depth-first-search

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func find(node *TreeNode, sum int, seen map[int]bool) bool {
    if node == nil {
        return false
    }

    if _, have := seen[sum-node.Val]; have {
        return true
    }

    seen[node.Val] = true

    return find(node.Left, sum, seen) ||
           find(node.Right, sum, seen)
}

func findTarget(root *TreeNode, k int) bool {
    return find(root, k, map[int]bool{})
}
