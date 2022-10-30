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

func walk(node *TreeNode) (int, int) {
    if node == nil {
        return 0, 0
    }

    if node.Left == nil && node.Right == nil {
        return 1, node.Val
    }

    ldepth, lsum := walk(node.Left)
    rdepth, rsum := walk(node.Right)
    sum := 0

    if ldepth < rdepth {
        sum += rsum
    } else if rdepth < ldepth {
        sum += lsum
    } else {
        sum += lsum + rsum
    }

    return 1+max(ldepth, rdepth), sum
}

func deepestLeavesSum(root *TreeNode) int {
    _, sum := walk(root)

    return sum
}
