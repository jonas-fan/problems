// #tree #depth-first-search #stack

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

/*
func preorder(root *Node) []int {
    out := []int{}
    stack := []*Node{root}

    for len(stack) > 0 {
        top := len(stack) - 1
        node := stack[top]
        stack = stack[:top]

        if node != nil {
            out = append(out, node.Val)

            for i := len(node.Children) - 1; i >= 0; i-- {
                stack = append(stack, node.Children[i])
            }
        }
    }

    return out
}
*/
