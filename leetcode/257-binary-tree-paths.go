/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func concat(base string, num int) string {
    if base == "" {
        return fmt.Sprintf("%d", num)
    }

    return fmt.Sprintf("%s->%d", base, num)
}

func walk(node *TreeNode, base string, out []string) []string {
    if node == nil {
        return out
    }

    base = concat(base, node.Val)

    if node.Left == nil && node.Right == nil {
        return append(out, base)
    }

    out = walk(node.Left, base, out)
    out = walk(node.Right, base, out)

    return out
}

func binaryTreePaths(root *TreeNode) []string {
    return walk(root, "", []string{})
}
