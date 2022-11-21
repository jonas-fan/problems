// #tree #binary-tree #binary-search-tree #depth-first-search

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insert(node *TreeNode, val int) *TreeNode {
    if node == nil {
        return &TreeNode{Val: val}
    }

    if node.Val < val {
        node.Right = insert(node.Right, val)
    } else if node.Val > val {
        node.Left = insert(node.Left, val)
    }

    return node
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    return insert(root, val)
}
