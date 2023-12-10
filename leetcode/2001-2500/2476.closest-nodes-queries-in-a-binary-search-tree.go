// #tree #binary-tree #binary-search-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min(lhs int, rhs int) int {
    if lhs < rhs {
        return lhs
    }

    return rhs
}

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func find(node *TreeNode, query int, out []int) []int {
    if node == nil {
        return out
    }

    if node.Val == query {
        out[0] = query
        out[1] = query
        return out
    }

    if node.Val < query {
        out[0] = node.Val
        return find(node.Right, query, out)
    } else if node.Val > query {
        out[1] = node.Val
        return find(node.Left, query, out)
    }

    return out
}

func closestNodes(root *TreeNode, queries []int) [][]int {
    out := make([][]int, len(queries))

    for i, query := range queries {
        out[i] = find(root, query, []int{-1, -1})
    }

    return out
}
