/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func isEven(num int) bool {
    return num & 0x1 == 0
}

func walk(node *TreeNode, evenParent bool, evenGrandParent bool) int {
    if node == nil {
        return 0
    }

    even := isEven(node.Val)
    sum := walk(node.Left, even, evenParent) +
           walk(node.Right, even, evenParent)

    if evenGrandParent {
        sum += node.Val
    }

    return  sum
}

func sumEvenGrandparent(root *TreeNode) int {
    return walk(root, false, false)
}
