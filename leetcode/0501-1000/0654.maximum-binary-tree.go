/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func makeMaxTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    bound := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] > nums[bound] {
            bound = i
        }
    }

    return &TreeNode{
        Val:   nums[bound],
        Left:  makeMaxTree(nums[0:bound]),
        Right: makeMaxTree(nums[bound+1:]),
    }
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
    return makeMaxTree(nums)
}
