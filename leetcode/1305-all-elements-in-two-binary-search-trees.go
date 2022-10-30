// #tree #binary-tree #binary-search-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func merge(lhs []int, rhs []int) []int {
    out := make([]int, 0, len(lhs)+len(rhs))

    for len(lhs) > 0 || len(rhs) > 0 {
        list := &lhs

        if len(lhs) > 0 && len(rhs) > 0 {
            if rhs[0] < lhs[0] {
                list = &rhs
            }
        } else if len(rhs) > 0 {
            list = &rhs
        }

        out = append(out, (*list)[0])
        (*list) = (*list)[1:]
    }

    return out
}

func walk(node *TreeNode, out []int) []int {
    if node == nil {
        return out
    }

    out = walk(node.Left, out)
    out = append(out, node.Val)
    out = walk(node.Right, out)

    return out
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    return merge(walk(root1, []int{}), walk(root2, []int{}))
}
