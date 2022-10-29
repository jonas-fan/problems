/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(node *TreeNode, out map[int]int) {
    if node == nil {
        return
    }

    out[node.Val]++

    dfs(node.Left, out)
    dfs(node.Right, out)
}

func findTarget(root *TreeNode, k int) bool {
    seen := map[int]int{}

    dfs(root, seen)

    for num := range seen {
        target := k - num

        if count, ok := seen[target]; !ok {
            continue
        } else if num == target && count < 2 {
            continue
        }

        return true
    }

    return false
}
