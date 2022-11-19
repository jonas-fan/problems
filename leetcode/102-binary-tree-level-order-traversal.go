// #tree #binary-tree #breadth-first-search

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

    out := [][]int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        values := []int{}
        nodes := []*TreeNode{}

        for _, node := range queue {
            values = append(values, node.Val)

            if node.Left != nil {
                nodes = append(nodes, node.Left)
            }

            if node.Right != nil {
                nodes = append(nodes, node.Right)
            }
        }

        out = append(out, values)
        queue = nodes
    }

    return out    
}
