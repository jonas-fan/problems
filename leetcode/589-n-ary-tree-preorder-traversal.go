/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func walk(node *Node, out []int) []int {
    if node == nil {
        return out
    }

    out = append(out, node.Val)

    for _, child := range node.Children {
        out = walk(child, out)
    }

    return out
}

func preorder(root *Node) []int {
    return walk(root, []int{})
}
