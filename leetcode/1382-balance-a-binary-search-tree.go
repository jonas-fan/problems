/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorder(node *TreeNode, out []int) []int {
    if node == nil {
        return out
    }

    out = inorder(node.Left, out)
    out = append(out, node.Val)
    out = inorder(node.Right, out)

    return out
}

func makeTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    index := len(nums) >> 1

    return &TreeNode{
        Val:   nums[index],
        Left:  makeTree(nums[:index]),
        Right: makeTree(nums[index+1:]),
    }
}

func balanceBST(root *TreeNode) *TreeNode {
    return makeTree(inorder(root, []int{}))
}
