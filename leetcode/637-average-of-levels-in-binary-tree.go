/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sum(nums []int) int {
    out := 0

    for _, num := range nums {
        out += num
    }

    return out
}

func dfs(node *TreeNode, level int, out [][]int) [][]int {
    if node == nil {
        return out
    }

    if level < len(out) {
        out[level] = append(out[level], node.Val)
    } else {
        out = append(out, []int{node.Val})
    }

    out = dfs(node.Left, level+1, out)
    out = dfs(node.Right, level+1, out)

    return out
}

func averageOfLevels(root *TreeNode) []float64 {
    buckets := dfs(root, 0, [][]int{})
    out := make([]float64, 0, len(buckets))

    for _, bucket := range buckets {
        average := float64(sum(bucket)) / float64(len(bucket))

        out = append(out, average)
    }

    return out
}
