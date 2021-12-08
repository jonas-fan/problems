/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findTilt(root *TreeNode) int {
    _, tilt := sum(root)

    return tilt
}

func sum(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }

    lsum, ltilt := sum(root.Left)
    rsum, rtilt := sum(root.Right)

    return root.Val + lsum + rsum, abs(lsum - rsum) + ltilt + rtilt
}

func abs(value int) int {
    if value < 0 {
        return -value
    }

    return value
}
