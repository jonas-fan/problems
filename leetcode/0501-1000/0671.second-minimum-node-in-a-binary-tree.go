/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func walk(node *TreeNode) int {
    if node == nil {
        return -1
    } else if node.Left == nil || node.Right == nil {
        return -1
    }

    lhs, rhs := -1, -1

    if node.Val == node.Left.Val {
        lhs = walk(node.Left)
    } else {
        lhs = node.Left.Val
    }

    if node.Val == node.Right.Val {
        rhs = walk(node.Right)
    } else {
        rhs = node.Right.Val
    }

    if lhs == -1 || rhs == -1 {
        return max(lhs, rhs)
    }

    return min(lhs, rhs)
}

func findSecondMinimumValue(root *TreeNode) int {
    return walk(root)
}
