/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func walk(node *TreeNode, subNode *TreeNode, matching bool) bool {
    if node == nil && subNode == nil {
        return matching
    } else if node == nil || subNode == nil {
        return false
    }

    if !matching {
        if node.Val == subNode.Val &&
           walk(node.Left, subNode.Left, true) &&
           walk(node.Right, subNode.Right, true) {
           return true 
        }

        return walk(node.Left, subNode, false) ||
               walk(node.Right, subNode, false)
    } else if node.Val == subNode.Val {
        return walk(node.Left, subNode.Left, true) &&
               walk(node.Right, subNode.Right, true)
    }

    return false
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    return walk(root, subRoot, false)
}
