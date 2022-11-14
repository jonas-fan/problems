// #tree #breadth-first-search #queue

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    out := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        values := []int{}
        nodes := []*TreeNode{}

        for _, node := range queue {
            if node == nil {
                continue
            }

            values = append(values, node.Val)
            nodes = append(nodes, node.Left)
            nodes = append(nodes, node.Right)
        }

        if len(values) > 0 {
            out = append(out, values)
        }

        queue = nodes
    }

    return out
}
