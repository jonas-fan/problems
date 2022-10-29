/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const (
    FALSE int = iota
    TRUE
    OR
    AND
)

func eval(node *TreeNode) bool {
    if node == nil {
        return false
    }

    // leaf
    if node.Left == nil && node.Right == nil {
        switch node.Val {
        case FALSE:
            return false
        default:
            return true
        }
    }

    switch node.Val {
    case OR:
        return eval(node.Left) || eval(node.Right)
    case AND:
        return eval(node.Left) && eval(node.Right)
    default:
        panic("Not supported operation")
    }
}

func evaluateTree(root *TreeNode) bool {
    return eval(root)
}
