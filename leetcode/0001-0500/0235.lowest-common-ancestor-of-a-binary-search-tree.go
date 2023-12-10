// #tree #binary-tree #binary-search-tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func walk(node *TreeNode, min *TreeNode, max *TreeNode) *TreeNode {
    if node == nil {
        return nil
    }

    if node.Val < min.Val {
        return walk(node.Right, min, max)
    } else if node.Val > max.Val {
        return walk(node.Left, min, max)
    }

    if node.Val == min.Val {
        return min
    } else if node.Val == max.Val {
        return max
    }

    return node
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p.Val > q.Val {
        p, q = q, p
    }

    return walk(root, p, q)
}
