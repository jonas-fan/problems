// #tree #binary-tree #binary-search-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }

    if root.Val < val {
        return searchBST(root.Right, val)
    } else if root.Val > val {
        return searchBST(root.Left, val)
    }

    return root
}