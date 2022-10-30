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

func walk(node *TreeNode, level int) (int, int) {
    if node == nil {
        return level, 0
    }

    depth := level + 1
    
    if node.Left == nil && node.Right == nil {
        return depth, node.Val
    }

    ldepth, lsum := walk(node.Left, depth)
    rdepth, rsum := walk(node.Right, depth)
    sum := 0

    if ldepth < rdepth {
        sum += rsum
    } else if rdepth < ldepth {
        sum += lsum
    } else {
        sum += lsum + rsum
    }

    return max(ldepth, rdepth), sum
}

func deepestLeavesSum(root *TreeNode) int {
    _, sum := walk(root, 0)

    return sum
}
