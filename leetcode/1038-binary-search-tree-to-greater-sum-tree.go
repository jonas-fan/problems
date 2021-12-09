/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bstToGst(root *TreeNode) *TreeNode {
    gst(root, 0)

    return root
}

func gst(root *TreeNode, base int) int {
    if root == nil {
        return 0
    }

    sum := gst(root.Right, base)
    val := root.Val

    root.Val += base + sum

    return val + sum + gst(root.Left, root.Val)
}
