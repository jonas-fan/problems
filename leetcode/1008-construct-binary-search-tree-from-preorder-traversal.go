// #tree #binary-tree #order

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func makeTree(preorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    } else if len(preorder) == 1 {
        return &TreeNode{Val: preorder[0]}        
    }

    value, preorder := preorder[0], preorder[1:]
    bound := 0

    for ; bound < len(preorder); bound++ {
        if preorder[bound] > value {
            break
        }
    }

    return &TreeNode{
        Val:   value,
        Left:  makeTree(preorder[:bound]),
        Right: makeTree(preorder[bound:]),
    }
}

func bstFromPreorder(preorder []int) *TreeNode {
    return makeTree(preorder)
}
