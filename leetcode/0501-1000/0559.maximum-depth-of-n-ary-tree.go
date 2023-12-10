/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func max(lhs int, rhs int) int {
    if lhs < rhs {
        return rhs
    }

    return lhs
}

func walk(node *Node, level int) int {
    if node == nil {
        return level
    }

    depth := level + 1

    for _, child := range node.Children {
        depth = max(depth, walk(child,level+1))
    }

    return depth
}

func maxDepth(root *Node) int {
    return walk(root, 0)
}
