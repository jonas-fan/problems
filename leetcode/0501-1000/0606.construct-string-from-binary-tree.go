/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func surround(str string) string {
    return fmt.Sprintf("(%s)", str)
}

func tree2str(root *TreeNode) string {
    if root == nil {
        return ""
    }

    out := fmt.Sprintf("%v", root.Val)

    if root.Left != nil || root.Right != nil {
        out += surround(tree2str(root.Left))

        if root.Right != nil {
            out += surround(tree2str(root.Right))
        }
    }

    return out
}
