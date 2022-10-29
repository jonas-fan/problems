/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    bound := len(nums) / 2

    return &TreeNode{
        Val:   nums[bound],
        Left:  dfs(nums[0:bound]),
        Right: dfs(nums[bound+1:]),
    }
}

func sortedArrayToBST(nums []int) *TreeNode {
    return dfs(nums)
}
