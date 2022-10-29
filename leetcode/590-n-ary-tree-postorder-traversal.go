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

    for _, child := range node.Children {
        out = walk(child, out)
    }

    out = append(out, node.Val)

    return out
}

func postorder(root *Node) []int {
    return walk(root, []int{})
}
