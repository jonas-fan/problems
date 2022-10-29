/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func eval(node *TreeNode) int {
    if node == nil {
        return 0
    }

    return node.Val
}

func lchild(node *TreeNode) *TreeNode {
    if node == nil {
        return nil
    }

    return node.Left
}

func rchild(node *TreeNode) *TreeNode {
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
        Val:   eval(root1) + eval(root2),
        Left:  mergeTrees(lchild(root1), lchild(root2)),
        Right: mergeTrees(rchild(root1), rchild(root2)),
    }
}
