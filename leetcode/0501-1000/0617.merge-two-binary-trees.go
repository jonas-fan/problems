// #tree #binary-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func value(node *TreeNode) int {
    if node == nil {
        return 0
    }

    return node.Val
}

func childLeft(node *TreeNode) *TreeNode {
    if node == nil {
        return nil
    }

    return node.Left
}

func childRight(node *TreeNode) *TreeNode {
    if node == nil {
        return nil
    }

    return node.Right
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil && root2 == nil {
        return nil
    }

    return &TreeNode{
        Val:   value(root1) + value(root2),
        Left:  mergeTrees(childLeft(root1), childLeft(root2)),
        Right: mergeTrees(childRight(root1), childRight(root2)),
    }
}
