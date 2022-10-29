/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func walk(node *TreeNode, parent *TreeNode, level int, num int) (int, *TreeNode) {
    if node == nil {
        return 0, nil
    }

    if node.Val == num {
        return level, parent
    }

    if ldepth, lparent := walk(node.Left, node, level+1, num); ldepth != 0 {
        return ldepth, lparent
    }

    return walk(node.Right, node, level+1, num)
}

func isCousins(root *TreeNode, x int, y int) bool {
    xdepth, xparent := walk(root, nil, 0, x)
    ydepth, yparent := walk(root, nil, 0, y)
    
    if xdepth != ydepth {
        return false
    }

    return xparent != yparent
}
