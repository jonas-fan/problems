/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// #tree #binary-tree #order

import (
    "sort"
)

func clone(nums []int) []int {
    out := make([]int, len(nums))

    copy(out, nums)

    return out
}

func makeTree(inorder []int, preorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }

    value := preorder[0]
    bound := 0

    for i := 0; i < len(inorder); i++ {
        if inorder[i] == value {
            bound = i
            break
        }
    }

    return &TreeNode{
        Val:   value,
        Left:  makeTree(inorder[:bound], preorder[1:1+bound]),
        Right: makeTree(inorder[bound+1:], preorder[1+bound:]),
    }
}

func bstFromPreorder(preorder []int) *TreeNode {
    inorder := clone(preorder)

    sort.Ints(inorder)

    return makeTree(inorder, preorder)
}
